package main

import (
	// "bufio"
	"fmt"
	"io/ioutil"
	"os/exec"
	// "time"
)

func main() {
	cmd := exec.Command("/bin/sh", "-c", "ping 127.0.0.1")
	// _, err := cmd.Output()
	// if err != nil {
	// 	panic(err.Error())
	// }

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("StdoutPipe: " + err.Error())
		return
	}

	if err := cmd.Start(); err != nil {
		panic(err.Error())
	}
	// go func() {
	// 	err := cmd.Run()
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 	}
	// }()

	// for {
	// 	r := bufio.NewReader(stdout)

	// 	line, _, err := r.ReadLine()
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 	}

	// 	fmt.Println("begin")
	// 	fmt.Println("stdout: %s", string(line))
	// }

	for {
		bytes, err := ioutil.ReadAll(stdout)
		if err != nil {
			fmt.Println("ReadAll stdout: ", err.Error())
			return
		}

		fmt.Println("stdout: %s", string(bytes))
	}

	// ticker := time.NewTicker(time.Second * 1)

	// go func() {
	// 	for t := range ticker.C {
	// 		fmt.Println("Ticker at", t)

	// 		info, _ := cmd.Output()
	// 		fmt.Println(string(info))

	// 		// bytes, err := ioutil.ReadAll(stdout)
	// 		// if err != nil {
	// 		// 	fmt.Println("ReadAll stdout: ", err.Error())
	// 		// 	return
	// 		// }

	// 		// fmt.Println("stdout: %s", string(bytes))
	// 	}
	// }()
	if err := cmd.Wait(); err != nil {
		panic(err.Error())
	}
	// time.Sleep(time.Second * 200)
	// ticker.Stop()
	// fmt.Println("Ticker stopped")
}
