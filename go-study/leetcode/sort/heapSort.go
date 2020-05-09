package main

import (
	"fmt"
)

func main() {
	nums := heapSort([]int{5, 1, 1, 2, 0, 0})
	fmt.Println(nums)
}

func heapSort(nums []int) []int {
	n := len(nums)
	for i := n/2-1 ; i >= 0; i-- {
		minFixDown(i, n, nums)
	}

	for i := n-1; i>0; i-- {
		maxHeapDeleteNumber(nums, i)
	}

	return nums
}

func minFixDown(i, n int, nums []int) {
	bigChild := 2*i+1
	tmp := nums[i]
	for bigChild < n {
		
		if bigChild + 1 < n && nums[bigChild] < nums[bigChild+1] {
			bigChild = bigChild +1
		}

		if nums[bigChild] <= tmp {
			break
		}

		nums[i] = nums[bigChild]
		i = bigChild
		bigChild = 2*i +1
	}

	nums[i] = tmp
}

func maxHeapDeleteNumber(nums []int, n int) {
	tmp := nums[0]
	nums[0] = nums[n]
	nums[n] = tmp
	minFixDown(0, n, nums)
}
