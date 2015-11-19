package main

import (
	"fmt"
)

// 插入排序算法
func insertionSort(value []float64) {
	for i := 1; i < len(value); i++ {
		key := value[i]
		j := i - 1
		for j >= 0 && value[j] > key {
			value[j+1] = value[j]
			j = j - 1
		}
		value[j+1] = key
	}
}

func main() {
	floatList := []float64{1.2, -1.1, 0, 1.0, 5.9, 5.90, -10, -89, 2, 0.00001, 0.00}
	fmt.Println("UnSort: ", floatList) // [1.2 -1.1 0 1 5.9 5.9 -10 -89 2 1e-05 0]
	insertionSort(floatList)
	fmt.Println("Sort: ", floatList) // [-89 -10 -1.1 0 0 1e-05 1 1.2 2 5.9 5.9]
}
