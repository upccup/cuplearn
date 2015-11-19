package main

import (
	"fmt"
)

func mergeSort(value []float64) []float64 {
	if len(value) <= 1 {
		return value
	}

	mid := (len(value) / 2)
	left := mergeSort(value[:mid])
	right := mergeSort(value[mid:])

	return merge(left, right)
}

func merge(left, right []float64) []float64 {
	var result []float64
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	result = append(result, left[1:]...)
	result = append(result, right[r:]...)
	return result

}

func main() {
	floatList := []float64{1.2, -1.1, 0, 1.0, 5.9, 5.90, -10, -89, 2, 0.00001, 0.00}
	fmt.Println("UnSort: ", floatList) // [1.2 -1.1 0 1 5.9 5.9 -10 -89 2 1e-05 0]
	floatList = mergeSort(floatList)
	fmt.Println("Sort: ", floatList) // [-89 -1.1 0 0 1 5.9]
}
