package main

import (
	"fmt"
	"os"
)

// func (f *File) Write(b []byte) (n int, err error)
// 参数列表: b 要写入的类容
// 返回值: 返回int 返回写入的字节数, 返回err 返回写入的错误对象
// 功能说明: 这个函数主要是往一个文件里写入类容

func main() {
	b := make([]byte, 10)
	fi, err := os.OpenFile("test.go", os.O_RDWR|os.O_APPEND, 0420)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	defer fi.Close()
	n, _ := fi.Read(b)
	fmt.Printf("len: %d, file content: %s\n", n, b[:n])
	fi.Write([]byte("cc\n"))
	fi.Seek(0, 0)
	n, _ = fi.Read(b)
	fmt.Printf("now len: %d, file content: %s\n", n, b[:n])
}

/*
len: 0, file content:
now len: 3, file content: cc
*/
