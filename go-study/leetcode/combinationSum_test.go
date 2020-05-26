package leetcode

import (
	"fmt"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	// result := combinationSum([]int{2, 3, 6, 7}, 7)
	// fmt.Println(result)
	result := combinationSum([]int{2, 3, 5}, 8)
	fmt.Println(result)
	t.Log(result)
}
