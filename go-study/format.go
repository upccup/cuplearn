package main

import (
	"fmt"
	"time"
)

func main() {
	s := fmt.Sprintf("aaaaa %s", "bbb")
	fmt.Println(s)
	fmt.Println(string(11))

	testChannel := make(chan string, 10)

	go func() {
		for i := 0; i < 90; i++ {
			testChannel <- "111"
			time.Sleep(time.Second * 3)
		}
	}()

	for test := range testChannel {
		fmt.Println(test)
	}
}
