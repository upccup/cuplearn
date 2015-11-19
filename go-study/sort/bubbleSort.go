package main

import (
	"fmt"
)

// 冒泡排序算法

func bubbleSort(value []float64) {
	for i := 0; i < len(value)-1; i++ {
		for j := i + 1; j < len(value); j++ {
			if value[i] > value[j] {
				value[i], value[j] = value[j], value[i]
			}
		}
	}
}

func main() {
	floatList := []float64{1.2, -1.1, 0, 1.0, 5.9, 5.90, -10, -89, 2, 0.00001, 0.00}
	fmt.Println("UnSort: ", floatList) // [1.2 -1.1 0 1 5.9 5.9 -10 -89 2 1e-05 0]
	bubbleSort(floatList)
	fmt.Println("Sort: ", floatList) // [-89 -10 -1.1 0 0 1e-05 1 1.2 2 5.9 5.9]
}
