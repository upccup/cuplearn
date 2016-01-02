package main

import (
	"log"
)

func main() {
	conn := Open()
	defer conn.Close()
	result, err := conn.Do("DECR", "mykey")
	if err != nil {
		log.Panicln(err)
	}

	log.Println(result)

}
