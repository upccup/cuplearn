package main

import (
	"fmt"
)

func changeArray(a [4]int) {
	a[0] = 999

	return
}

func changeSlice(a []int) {
	a[0] = 998

	a = append(a, []int{9, 99}...)
	return
}

func changeSlice2(a *[]int) {
	(*a)[0] = 999

	*a = append(*a, []int{9, 99}...)
	return
}

func main() {
	var array [4]int = [4]int{1, 2, 3, 4}
	fmt.Println("array before change: ", array)

	changeArray(array)
	fmt.Println("array after change: ", array)

	var slice []int = []int{1, 2, 3, 4}
	fmt.Println("slice before change: ", slice)

	changeSlice(slice)
	fmt.Println("slice after change value: ", slice)

	changeSlice2(&slice)
	fmt.Println("slice after change point: ", slice)

}

/*
array before change:  [1 2 3 4]
array after change:  [1 2 3 4]
slice before change:  [1 2 3 4]
slice after change value:  [998 2 3 4]
slice after change point:  [999 2 3 4 9 99]
*/
