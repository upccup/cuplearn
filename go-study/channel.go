package main

import (
	"fmt"
	"time"
)

func Count(ch chan int) {
	fmt.Println(ch)
	ch <- 1
	fmt.Println("Counting")
}

func testerr(err chan<- error) {
	err <- nil

}

func main() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}
	for i, ch := range chs {
		//	value := <-ch
		//	fmt.Println(value)
		fmt.Println(i)

		<-ch
	}

	res := make(chan error, 10)

	for i := 0; i < 20; i++ {
		go func() {
			time.Sleep(time.Second * 1)
			testerr(res)
		}()

	}

	for {
		select {
		case err := <-res:
			fmt.Println(err)
		default:

		}
	}
}
