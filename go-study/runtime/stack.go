package main

import (
	"fmt"
	"runtime"
)

func echo(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	buffer := make([]byte, 4096)
	buffer = buffer[:runtime.Stack(buffer, true)]

	go echo("bye")

	go echo("world")

	echo("good")

	fmt.Println("# go routines : ", runtime.NumGoroutine())

	fmt.Println("Stack trace bytes : ")
	fmt.Println(string(buffer))

	buffer = buffer[:runtime.Stack(buffer, true)]
	fmt.Println(string(buffer))
}

/*
good
world
bye
good
world
bye
good
world
bye
good
world
bye
good
# go routines :  6
Stack trace bytes :
goroutine 1 [running]:
main.main()
	/Users/yaoyun/go/src/gotest/git-study/cuplearn/go-study/runtime/stack.go:19 +0x90

goroutine 1 [running]:
main.main()
	/Users/yaoyun/go/src/gotest/git-study/cuplearn/go-study/runtime/stack.go:32 +0x475
*/
