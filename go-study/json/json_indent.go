package main

import (
	// "encoding/json"
	"fmt"
	"time"
)

func main() {
	fmt.Println("aaaaaa")
	test := time.NewTicker(time.Second * 3)
	for {
		select {
		case <-test.C:
			fmt.Println("aaaaaaaaaaaaaaaaabbbb")
			for i := 0; i < 5; i++ {
				time.Sleep(time.Second * 1)
				fmt.Println(i)
			}

		}
	}
}
