package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://192.168.111.119:5051/state.json")

	if err != nil {
		fmt.Println(err)
		return
	}

	fi, err := os.OpenFile("test.go", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0420)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer fi.Close()
	robots, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fi.Write(robots)
	fmt.Println("over")

}
