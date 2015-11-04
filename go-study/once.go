package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

func main() {
	for i, v := range make([]string, 10) {
		once.Do(onces)
		fmt.Println("count:", v, "----", i)
	}

	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onced)
			fmt.Println("123")
		}()
	}

	time.Sleep(time.Second * 2)
}

func onces() {
	fmt.Println("onces")
	time.Sleep(time.Second * 5)
}

func onced() {
	fmt.Println("onced")
}
