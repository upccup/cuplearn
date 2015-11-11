package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	files, _ := ioutil.ReadDir("../")
	for _, f := range files {
		fmt.Println("Name", f.Name())
		fmt.Println("ModiTime", f.ModTime())
		fmt.Println("IsDir", f.IsDir())
	}
}
