package main

import (
	"fmt"
)

func main() {
	channel := make(chan int, 20)

	var i int
	for i = 0; i <= 10; i++ {
		channel <- i
	}

	fmt.Println(len(channel))

}
