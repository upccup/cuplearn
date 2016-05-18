package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s int64
	s = 12
	fmt.Println(s) //12
	if s, err := strconv.ParseInt("23", 10, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s) // int64 23
	} else {
		fmt.Println(err)
	}
	fmt.Println(s) //12
	s = 25
	fmt.Println(s) // 25

	var err error
	if s, err = strconv.ParseInt("23", 10, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s) // int64 23
	} else {
		fmt.Println(err)
	}
	fmt.Println(s) // 23

	//s := 11  build error
}
