package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("begin")

	timeout := make(chan bool, 1)

	go func() {
		time.Sleep(1e9)
		timeout <- true
	}()
	ch := make(chan int)

	select {
	case <-ch:
		fmt.Println("aaa")
	case <-timeout:
		fmt.Println("timeout")

	}

	fmt.Println("end")

}
