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
	// ten minute after:  2015-12-09 22:44:15.140701673 +0800 CST
	fmt.Println("ten minute ago: ", now.Add(time.Minute*-10))
	// ten minute ago:  2015-12-09 22:24:15.140701673 +0800 CST
}
