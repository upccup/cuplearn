package main

import (
	"log"
)

func add1(x int) int {
	x = x + 1
	return x
}

func add2(x *int) int {
	*x = *x + 2
	return *x
}

func main() {
	var x = 1

	log.Println("x = ", x)

	x1 := add1(x)
	log.Println("x+1 = ", x1)
	log.Println("x = ", x)

	x2 := add2(&x)
	log.Println("x+2 = ", x2)
	log.Println("x = ", x)
}

// out put
/*
x =  1
x+1 =  2
x =  1
x+2 =  3
x =  3
*/
