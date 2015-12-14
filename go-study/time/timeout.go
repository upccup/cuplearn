package main

import (
	"fmt"
	"time"
)

func main() {
	ontime := make(chan string)
	for {
		fmt.Println(timeoutTest(ontime))
	}

	// time.Sleep(time.Second * 5)
}

func timeoutTest(ontime chan string) string {

	go timeout(ontime)

	for {
		select {
		case ontimeString := <-ontime:
			return ontimeString
		case <-time.After(time.Second * 3):
			return "time out"
		}
	}
}

func timeout(ontime chan string) {
	time.Sleep(time.Second * 5)
	fmt.Println("execute")
	ontime <- "on time"
	return
}
