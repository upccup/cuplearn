package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	body := req.Body
	defer body.Close()
	x, _ := ioutil.ReadAll(body)
	fmt.Println(string(x))
	io.WriteString(w, string(x))
}

func main() {
	http.HandleFunc("/tesr", HelloServer)
	http.HandleFunc("/test", healthCheckHandle)
	go http.ListenAndServe(":12345", nil)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func healthCheckHandle(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "agent http is running good")
}
