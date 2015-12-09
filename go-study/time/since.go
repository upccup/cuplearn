package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

// func Since(t Time) Duration
// 功能说明: 返回时间t到当前时间的间隔, time.Now().Sub(t)的简写

func main() {
	t := time.Now()
	time.Sleep(time.Second * 3)
	fmt.Println(time.Since(t)) // 3.00505111s

	f := time.Now().Add(10 * time.Minute)
	fmt.Println(time.Since(f)) // -9m59.999999805s

	files, _ := ioutil.ReadDir("../")
	for _, f := range files {
		fmt.Printf("File %s modify %s ageo\n", f.Name(), time.Since(f.ModTime()))
		// File channel.go modify 1865h45m22.540498739s ageo
		// ...
	}
}
