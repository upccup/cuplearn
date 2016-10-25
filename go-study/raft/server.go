package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/goraft/raft"
	"github.com/gorilla/mux"
)

var verbose bool
var trace bool
var debug bool
var host string
var port int
var join string

func init() {
	flag.BoolVar(&verbose, "v", false, "verbose logging")
	flag.BoolVar(&trace, "trace", false, "Raft trace debugging")
	flag.BoolVar(&debug, "debug", false, "Raft debugging")
	flag.StringVar(&host, "h", "localhost", "hostname")
	flag.IntVar(&port, "p", 4001, "port")
	flag.StringVar(&join, "join", "", "host:port of leader to join")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [arguments] <data-path> \n", os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {
	log.SetFlags(0)
	flag.Parse()

	if verbose {
		log.Print("Verbose logging enabled.")
	}

	if trace {
		raft.SetLogLevel(raft.Trace)
		log.Print("Raft trace debugging enabled.")
	} else if debug {
		raft.SetLogLevel(raft.Debug)
		log.Print("Raft debugging enabled.")
	}

	rand.Seed(time.Now().UnixNano())

	raft.RegisterCommand(&WriteCommand{})

	if flag.NArg() == 0 {
		flag.Usage()

		log.Fatal("Data path argument required")
	}

	path := flag.Arg(0)

	if err := os.MkdirAll(path, 0744); err != nil {
		log.Fatalf("Unable to creater path: %s", err.Error())
	}

	log.SetFlags(log.LstdFlags)
	s := NewServer(path, host, port)
	log.Fatal(s.ListenAndServer(join))
}

type Server struct {
	name       string
	host       string
	port       int
	path       string
	router     *mux.Router
	raftServer raft.Server
	httpServer *http.Server
	db         *DB
	mutex      sync.RWMutex
}

func NewServer(path, host string, port int) *Server {
	s := &Server{
		host:   host,
		port:   port,
		path:   path,
		db:     NewDb(),
		router: mux.NewRouter(),
	}

	if b, err := ioutil.ReadFile(filepath.Join(path, "name")); err == nil {
		s.name = string(b)
	} else {
		s.name = fmt.Sprintf("%07x", rand.Int())[0:7]
		if err = ioutil.WriteFile(filepath.Join(path, "name"), []byte(s.name), 0644); err != nil {
			log.Fatal(err)
		}
	}

	return s
}

func (s *Server) connectionString() string {
	return fmt.Sprintf("http://%s:%d", s.host, s.port)
}

func (s *Server) ListenAndServer(leader string) error {
	var err error

	log.Printf("Initializing Raft Server :%s", s.path)

	transporter := raft.NewHTTPTransporter("/raft", 200*time.Millisecond)
	s.raftServer, err = raft.NewServer(s.name, s.path, transporter, nil, s.db, "")
	if err != nil {
		log.Fatal(err)
	}

	transporter.Install(s.raftServer, s)
	s.raftServer.Start()

	if leader != "" {
		log.Println("Attempting to join leader: ", leader)

		if !s.raftServer.IsLogEmpty() {
			log.Fatal("Cannot join with existing log")
		}

		if err := s.Join(leader); err != nil {
			log.Fatal(err)
		}
	} else if s.raftServer.IsLogEmpty() {
		log.Println("Initializing new cluster")

		_, err := s.raftServer.Do(&raft.DefaultJoinCommand{
			Name:             s.raftServer.Name(),
			ConnectionString: s.connectionString(),
		})

		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Println("Recovered from log")
	}

	log.Println("Initializing HTTP server")

	s.httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: s.router,
	}

	s.router.HandleFunc("/db/{key}", s.readHandler).Methods("GET")
	s.router.HandleFunc("/leader", s.leaderHandler).Methods("GET")
	s.router.HandleFunc("/state", s.stateHandler).Methods("GET")
	s.router.HandleFunc("/db/{key}", s.writeHandler).Methods("POST")
	s.router.HandleFunc("/join", s.joinHandler).Methods("POST")

	log.Println("Listening at: ", s.connectionString)

	return s.httpServer.ListenAndServe()
}

func (s *Server) HandleFunc(pattern string, Handler func(http.ResponseWriter, *http.Request)) {
	s.router.HandleFunc(pattern, Handler)
}

func (s *Server) Join(leader string) error {
	command := &raft.DefaultJoinCommand{
		Name:             s.raftServer.Name(),
		ConnectionString: s.connectionString(),
	}

	var b bytes.Buffer
	json.NewEncoder(&b).Encode(command)
	resp, err := http.Post(fmt.Sprintf("http://%s/join", leader), "application/json", &b)
	if err != nil {
		return err
	}

	resp.Body.Close()
	return nil
}

func (s *Server) joinHandler(w http.ResponseWriter, req *http.Request) {
	command := &raft.DefaultJoinCommand{}

	if err := json.NewDecoder(req.Body).Decode(&command); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := s.raftServer.Do(command); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	return
}

func (s *Server) readHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	value := s.db.Get(vars["key"])
	w.Write([]byte(value))
}

func (s *Server) writeHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	value := string(b)

	_, err = s.raftServer.Do(NewWriteCommand(vars["key"], value))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (s *Server) leaderHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(s.raftServer.State()))
}

func (s *Server) stateHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(s.raftServer.Leader()))
}
