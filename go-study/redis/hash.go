package main

import (
	"log"

	redis "github.com/garyburd/redigo/redis"
	redisgo "gotest/git-study/cuplearn/go-study/redis/conn"
)

func WriteHashToRedis(key, field, value string, timeout int) error {
	conn := redisgo.Open()
	defer conn.Close()
	var err error
	log.Printf("redis HSET: %s, field: %s, value: %s", key, field, value)
	if _, err = conn.Do("HSET", key, field, value); err != nil {
		return err
	}

	if timeout != -1 {
		_, err = conn.Do("EXPIRE", key, timeout)
		return err
	}
	return nil
}

func GetHashValueFromRedis(key string) (string, error) {
	conn := redisgo.Open()
	defer conn.Close()
	status, err := redis.String(conn.Do("HGET", key, "status"))

	if err != nil {
		return "", err
	}

	return status, err
}

func main() {
	err := WriteHashToRedis("myhash", "field", "111111", -1)
	if err != nil {
		log.Panicln(err)
	}

	log.Println("WriteHashToRedis success")

	status, err := GetHashValueFromRedis("myhash")
	if err == redis.ErrNil {
		log.Println("status is nil")
	} else if err != nil {
		log.Panicln(err)
	}

	log.Println("GetHashValueFromRedis ", status)

}
