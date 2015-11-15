package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a chan int
	func(c <-chan int) { // <-chan 表示的是信道接受数据
		var chandir = reflect.TypeOf(c).ChanDir()
		fmt.Println(chandir == reflect.RecvDir) // reflect.RecvDir 常量表示是接受数据的信道
		// true
	}(a)

	func(c chan<- int) { // chan<-  代表信道发送数据
		var chandir = reflect.TypeOf(c).ChanDir()
		fmt.Println(chandir == reflect.SendDir) // reflect.SendDir 常量表示发送数据的信道
		// true
	}(a)

	func(c chan int) { // chan 表示是双向信道
		var chandir = reflect.TypeOf(c).ChanDir()
		fmt.Println(chandir == reflect.BothDir) // reflect.BothDir 常量表示的是双向信道
		// true
	}(a)

}
