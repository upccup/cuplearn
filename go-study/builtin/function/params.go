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

func multiArgs(name string, books ...string) {
	log.Printf("name: %s books: %+v", name, books)
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

	multiArgs("Aaaaa")
	multiArgs("Bbbbb", "C")
	multiArgs("Ccccc", "C", "C++")
	books := []string{"C", "C++", "C#"}
	multiArgs("Ddddd", books...)
}

// out put
/*
2016/10/31 11:13:00 x =  1
2016/10/31 11:13:00 x+1 =  2
2016/10/31 11:13:00 x =  1
2016/10/31 11:13:00 x+2 =  3
2016/10/31 11:13:00 x =  3
2016/10/31 11:13:00 name: Aaaaa books: []
2016/10/31 11:13:00 name: Bbbbb books: [C]
2016/10/31 11:13:00 name: Ccccc books: [C C++]
2016/10/31 11:13:00 name: Ddddd books: [C C++ C#]
*/
