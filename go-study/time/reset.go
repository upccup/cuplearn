package main

import (
	"log"
	"time"
)

func main() {
	timer := time.NewTicker(time.Second * 5)

	go func() {
		for {
			select {
			case <-timer.C:
				log.Println("abcdefg")
			}
		}
	}()

	timerReset := time.NewTicker(time.Second * 3)

	for {
		select {
		case <-timerReset.C:
			log.Println("sasd")
		}
	}

}
