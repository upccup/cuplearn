package main

import (
	"log"

	redis "github.com/garyburd/redigo/redis"
	redisgo "gotest/git-study/cuplearn/go-study/redis/conn"
)

func main() {
	conn := redisgo.Open()
	defer conn.Close()

	conn.Send("MULTI")
	conn.Send("GET", "count")
	conn.Send("INCR", "count")
	conn.Send("GET", "count")
	conn.Send("GET", "countA")

	r, err := conn.Do("EXEC")
	if err != nil {
		log.Panicln(err)
	}
	log.Println(r)

	status, err := redis.String(conn.Do("GET", "count"))
	if err != nil {
		log.Panic(err)
	}

	log.Println(status)
}
