package main

import (
	"fmt"
)

func main() {
	nums := []int{5, 0, 6, 3, 1, 2}
	quickSort(nums)
	fmt.Println(nums)
	nums = []int{5, 2, 3, 1}
	quickSort(nums)
	fmt.Println(nums)
}

func quickSort(nums []int) []int {
	sort(0, len(nums)-1, nums)
	return nums
}

func sort(left, right int, nums []int) {
	if left >= right {
		return
	}

	n := nums[left]
	i := left
	j := right

	for i < j {
		for nums[j] >= n && i < j {
			j--
		}

		for nums[i] <= n && i < j {
			i++
		}

		if i < j {
			swap(i, j, nums)
		}
	}

	swap(left, i, nums)
	sort(left, i-1, nums)
	sort(i+1, right, nums)

}

func swap(i, j int, nums []int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
	return
}
