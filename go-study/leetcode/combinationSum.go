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

// https://leetcode.com/problems/combination-sum-iii/
func combinationSumIII(k, n int) [][]int {
	candidates := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := [][]int{}
	combineIII(candidates, k, n, []int{}, &result)
	return result
}

func combineIII(candidates []int, k int, target int, curComb []int, result *[][]int) {
	if target == 0 && len(curComb) == k {
		*result = append(*result, append([]int{}, curComb...))
		return
	} else if target < 0 {
		return
	} else {
		for i, v := range candidates {
			combineIII(candidates[i+1:], k, target-v, append(curComb, v), result)
		}
	}
}

// 题目描述 https://leetcode.com/problems/combination-sum-iv/
// 方案参考: https://leetcode.com/problems/combination-sum-iv/discuss/433597/Go-0ms-double-100-solution
// 思路解析: https://github.com/grandyang/leetcode/issues/377
func combinationSumIV(nums []int, target int) int {
	sort.Ints(nums)
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 0; i <= target; i++ {
		for _, v := range nums {
			if i+v > target {
				break
			}

			dp[i+v] += dp[i]
		}
	}
	return dp[target]
}
