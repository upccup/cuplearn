package main

import (
	"fmt"
)

func shellsort(value []float64) {
	h := 1
	for h < len(value) {
		h = h*3 + 1
	}

	for h > 1 {
		for i := h; i < len(value); i++ {
			for j := i; j >= h && value[j] < value[h]; j = j - h {
				value[j], value[j-h] = value[j-h], value[j]
			}
		}
		h = h / 3
	}
}

func main() {
	floatList := []float64{1.2, -1.1, 0, 1.0, 5.9, 5.90, -10, -89, 2, 0.00001, 0.00}
	fmt.Println("UnSort: ", floatList) // [1.2 -1.1 0 1 5.9 5.9 -10 -89 2 1e-05 0]
	shellsort(floatList)
	fmt.Println("Sort: ", floatList) // [-89 -10 -1.1 0 0 1e-05 1 1.2 2 5.9 5.9]
}
