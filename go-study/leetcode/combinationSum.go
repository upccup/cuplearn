package leetcode

import (
	"reflect"
	"sort"
)

var result [][]int

// https://leetcode.com/problems/combination-sum/
func combinationSum(candidates []int, target int) [][]int {
	return combine(0, target, []int{}, candidates)
}

func combine(sum int, target int, curComb []int, candidates []int) [][]int {
	var tmp, result [][]int
	if sum == target {
		// return [][]int{curComb} 这么写会有一个问题当 curComb值变化时, result值也会有变化,因为这里两个其实指向了同一个底层数组
		return [][]int{append([]int{}, curComb...)}
	} else if sum < target {
		for i, v := range candidates {
			tmp = combine(sum+v, target, append(curComb, v), candidates[i:])
			result = append(result, tmp...)
		}
	}
	return result
}

// https://leetcode.com/problems/combination-sum-ii/
func combinationSumII(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return combineII(0, target, []int{}, candidates)
}

func combineII(sum int, target int, curComb []int, candidates []int) [][]int {
	var tmp, result [][]int
	if sum == target {
		return [][]int{append([]int{}, curComb...)}
	} else if sum < target {
		for i, v := range candidates {
			tmp = combineII(sum+v, target, append(curComb, v), candidates[i+1:])
			result = merge(result, tmp)
		}
	}

	return result
}

func merge(a, b [][]int) [][]int {
	for _, v := range b {
		if !contain(a, v) {
			a = append(a, v)
		}
	}

	return a
}

func contain(a [][]int, b []int) bool {
	for _, v := range a {
		if reflect.DeepEqual(v, b) {
			return true
		}
	}

	return false
}

// https://leetcode.com/problems/combination-sum-ii/discuss/498085/Go-0ms
func combinationSumIIOptimize(candidates []int, target int) [][]int {
	result := [][]int{}
	sort.Ints(candidates)
	combineIIOptimize(candidates, target, []int{}, &result)
	return result
}

func combineIIOptimize(candidates []int, target int, curComb []int, result *[][]int) {
	if len(candidates) == 0 {
		return
	}

	if candidates[0] == target {
		curComb = append(curComb, candidates[0])
		*result = append(*result, curComb)
		return
	} else if candidates[0] > target {
		return
	} else {
		for i := 0; i < len(candidates)-1; i++ {
			// 去重
			if candidates[i] != candidates[i+1] {
				combineIIOptimize(candidates[i+1:], target, curComb, result)
				curComb2 := make([]int, len(curComb))
				copy(curComb2, curComb)
				curComb2 = append(curComb, candidates[0])
				combineIIOptimize(candidates[1:], target-candidates[0], curComb2, result)
				return
			}
		}

		curComb3 := make([]int, len(curComb))
		copy(curComb3, curComb)
		curComb3 = append(curComb3, candidates[0])
		combineIIOptimize(candidates[1:], target-candidates[0], curComb3, result)
		return
	}
}
