package main

import (
	"fmt"
	"os"
)

// func (f *File) Seek(offset int64, whence int) (ret int64, err error)
// 参数列表: offset 文件指针的位置 whence 相对位置标识
// 返回值: 返回ret 返回文件指针的位置, 返回 err 返回err错误对象
// 功能说明: 这个函数主要是把文件指针移动到文件制定的位置, whence为0时代表相对文件开始的位置
// 1 代表相对当前位置, 2 代表相对文件结尾的位置

func main() {
	b := make([]byte, 10)
	fi, err := os.Open("test.go")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()
	fi.Seek(3, 0)
	n, _ := fi.Read(b)
	fmt.Printf("%d\n", n)
	fmt.Printf("%s\n", b[:n])
}

/*
文件内容
cdcabbbbb
cc

10
abbbbb
cc
*/
