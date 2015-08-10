package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"
)

func cmdLs() {
	// arv := []string{"-c ls -a"}
	cmd := exec.Command("/bin/sh", "-c", "ls -a")
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

func cmdPwd() {
	var output bytes.Buffer
	cmd := exec.Command("pwd")
	cmd.Stdout = &output
	stdin, _ := cmd.StdinPipe()
	cmd.Start()
	stdin.Write([]byte("hello world"))
	stdin.Close()
	cmd.Wait()
	fmt.Printf("OUTPUT: %s\n", string(output.Bytes()))
}

func cmdCat() {
	var output bytes.Buffer

	cmd := exec.Command("cat")

	cmd.Stdout = &output

	stdin, _ := cmd.StdinPipe()
	cmd.Start() //执行开始
	stdin.Write([]byte("widuu test"))
	stdin.Close()
	cmd.Wait()                                        //等待执行完成
	fmt.Printf("The output is: %s\n", output.Bytes()) //The output is: widuu test!
}

func main() {
	cmdLs()
	cmdPwd()
	cmdCat()
}
