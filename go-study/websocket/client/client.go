package main

import (
	"flag"
	"log"
	"net/url"
	"time"

	// cli "github.com/codegangsta/cli"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:9009", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/echo"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial: ", err)
	}
	defer c.Close()

	go func() {
		defer c.Close()
		for {
			log.Println()
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read: ", err)
				break
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second * 3)
	defer ticker.Stop()

	for t := range ticker.C {
		err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
		if err != nil {
			log.Println("write: ", err)
			time.Sleep(time.Second * 3)
		}
	}

}
