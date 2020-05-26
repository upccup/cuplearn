package leetcode

import (
	"sort"
)

var result [][]int

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	return combine(0, target, []int{}, candidates)
}

func combine(sum int, target int, curComb []int, candidates []int) [][]int {
	var tmp, result [][]int
	if sum == target {
		// return [][]int{curComb} 这么写会有一个问题当 curComb值变化时, result值也会有变化,因为这里两个其实指向了同一个底层数组
		return [][]int{append([]int{}, curComb...)}
	} else if sum < target {
		for i, v := range candidates {
			// sum = sum + v
			// curComb = append(curComb, v)
			// tmp = combine(sum, target, curComb, candidates[i:])
			tmp = combine(sum+v, target, append(curComb, v), candidates[i:])
			result = append(result, tmp...)
		}
	}
	return result
}
