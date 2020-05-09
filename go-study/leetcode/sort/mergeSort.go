package main

import (
    "fmt"
)

func main() {
    nums := mergeSort([]int{5, 1, 1, 2, 0, 0}) 
    fmt.Println(nums)
}

func mergeSort(nums []int) []int {
    length := len(nums)
    if(length < 2) {
        return nums;
    }
    
    middle := length/2
    left := nums[0:middle]
    right := nums[middle:]

    return merge(mergeSort(left), mergeSort(right)) 
}

func merge(left, right []int) []int {
    leftLength := len(left)
    rightLength := len(right) 

    var result []int
    i := 0
    j := 0

    for i<leftLength && j < rightLength {
        if left[i] < right[j] {
            result = append(result, left[i])
            i++
        } else {
            result = append(result, right[j])
            j++
        }
    }

    for i < leftLength {
        result = append(result ,left[i])
        i++
    }

    for j < rightLength {
        result = append(result, right[j])
        j++
    }

    return result
}