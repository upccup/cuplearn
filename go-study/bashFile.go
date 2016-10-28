package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	cmd := exec.Command("/bin/sh", "-c", "bash ./atest.sh")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("StdoutPipe: " + err.Error())
		return
	}

	// stdin, err := cmd.StdinPipe()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// stdin.Write([]byte("hello world"))
	// stdin.Close()

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println("StderrPipe: ", err.Error())
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("Start: ", err.Error())
		return
	}

	bytesErr, err := ioutil.ReadAll(stderr)
	if err != nil {
		fmt.Println("ReadAll stderr: ", err.Error())
	}

	if len(bytesErr) != 0 {
		fmt.Println("stderr is not nil: %s", string(bytesErr))
		return
	}

	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("ReadAll stdout: ", err.Error())
		return
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("Wait: ", err.Error())
		return
	}

	fmt.Println("stdout: %s", string(bytes))
}
