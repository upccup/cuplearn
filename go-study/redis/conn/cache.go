package conn

import (
	"fmt"
	"log"
	"sync"

	redis "github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

const (
	cacheHost = "localhost"
	cachePort = 6379
	poolSize  = 100
)

func Open() redis.Conn {
	if pool != nil {
		return pool.Get()
	}

	mutex := &sync.Mutex{}
	mutex.Lock()
	InitCache()
	defer mutex.Unlock()

	return pool.Get()
}

func initConn() (redis.Conn, error) {
	addr := fmt.Sprintf("%s:%d", cacheHost, cachePort)
	c, err := redis.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	return c, err
}

func InitCache() {
	pool = redis.NewPool(initConn, poolSize)
	conn := Open()
	defer conn.Close()
	pong, err := conn.Do("ping")
	if err != nil {
		log.Panicln("can't connect cache server has error", err)
	}
	log.Println("reach cache server ", pong)
}

func DestroyCache() {
	log.Println("destroying Cache")
	if pool != nil {
		pool.Close()
		log.Println("cache was closed")
	}
}
