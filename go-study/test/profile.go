package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	// "os"
	// "runtime/pprof"
	"time"

	// "github.com/davecheney/profile"
)

func main() {
	// defer profile.Start(profile.MemProfile).Stop()
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	// fmt.Println("aaaaaa")
	// f, err := os.Create("test.prof")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// pprof.StartCPUProfile(f)
	// defer pprof.StopCPUProfile()

	time.Sleep(time.Second * 10)
	for i := 0; i < 999; i++ {
		go PrintFile()
		time.Sleep(time.Millisecond * 100)
	}
}

func PrintFile() {
	b, e := ioutil.ReadFile("./profile.go")
	if e != nil {
		fmt.Println("read file error")
		return
	}

	fmt.Println(string(b))
}
