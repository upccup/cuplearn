package main

import (
	"errors"
	"fmt"
)

func main() {
	var a int = 1
	var b *int = &a
	var c **int = &b
	var x int = *b
	fmt.Println("a = ", a)
	fmt.Println("&a = ", &a)
	fmt.Println("*&a = ", *&a)
	fmt.Println("b = ", b)
	fmt.Println("&b = ", &b)
	fmt.Println("*&b = ", *&b)
	fmt.Println("*b = ", *b)
	fmt.Println("c = ", c)
	fmt.Println("*c = ", *c)
	fmt.Println("&c = ", &c)
	fmt.Println("*&c = ", *&c)
	fmt.Println("**c = ", **c)
	fmt.Println("***&*&*&*&c = ", ***&*&*&*&c)
	fmt.Println("x = ", x)

	err := errors.New("error test")
	fmt.Println(err.Error())

	fmt.Println(test())
}

func test() string {
	var test string
	return test
}
