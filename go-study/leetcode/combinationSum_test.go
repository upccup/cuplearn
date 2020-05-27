package leetcode

import (
	"fmt"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	fmt.Println("TestCombinationSum: ")
	result := combinationSum([]int{2, 3, 6, 7}, 7)
	fmt.Println(result)
	result = combinationSum([]int{2, 3, 5}, 8)
	fmt.Println(result)
	t.Log(result)
}

func TestCombinationSumII(t *testing.T) {
	fmt.Println("TestCombinationSumII: ")
	result := combinationSumII([]int{10, 1, 2, 7, 6, 1, 5}, 8)
	fmt.Println(result)
	fmt.Println("TestCombinationSumIIOptimize: ")
	result = combinationSumIIOptimize([]int{10, 1, 2, 7, 6, 1, 5}, 8)
	fmt.Println(result)
}
