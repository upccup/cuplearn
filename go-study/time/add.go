package main

import (
	"fmt"
	"time"
)

// func (t Time) Add(d Duration) Time
//功能说明 返回时间(t+d)

func main() {
	now := time.Now()
	fmt.Println("now", now)
	fmt.Println("ten minute after: ", now.Add(time.Minute*10))
	fmt.Println("ten minute ago: ", now.Add(time.Minute*-10))
}
