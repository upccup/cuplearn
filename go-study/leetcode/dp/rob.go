package dp

// https://leetcode.com/problems/house-robber/
/**
使用动态规划解题最主要是找到状态转移方程,本题中状态转移方程为: dp[i] = max(dp[i-1], dp[i-2]+nums[i]), 即:
对当前i来说，有抢和不抢两种互斥的选择，不抢即为 dp[i-1]（等价于去掉 nums[i] 只抢 [0, i-1] 区间最大值），抢即为 dp[i-2] + nums[i]（等价于去掉 nums[i-1]）
参考文档: https://github.com/grandyang/leetcode/issues/198
**/
func rob(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	} else if length == 1 {
		return nums[0]
	} else if length == 2 {
		return max(nums[0], nums[1])
	}

	dp := make([]int, length)
	dp[0], dp[1] = nums[0], max(nums[0], nums[1])
	for i := 2; i <= length-1; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[length-1]
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}

/**
改进算法: 思想和上面的动态规划一样但是使用两个变量维护抢两个相邻目标的收益(刚好对应奇数和偶数)
在遍历数组时如果是抢劫偶数位置则robEven(偶数)+当前值,奇数位置则为robOdd加上当前值,这种分发可以保证组成最大的数字不相邻
最后去较大值返回
**/
func rob2(nums []int) int {
	var robEven, robOdd int
	for i := 1; i <= len(nums); i++ {
		if i%2 == 0 {
			robEven = max(robEven+nums[i-1], robOdd)
		} else {
			robOdd = max(robOdd+nums[i-1], robEven)
		}
	}

	return max(robEven, robOdd)
}

/**
改进算法2: 使用两个变量 rob 和 notRob，其中 rob 表示抢当前的房子，notRob 表示不抢当前的房子，
那么在遍历的过程中，先用两个变量 preRob 和 preNotRob 来分别记录更新之前的值，由于 rob 是要抢当前的房子，
那么前一个房子一定不能抢，所以使用 preNotRob 加上当前的数字赋给 rob，
然后 notRob 表示不能抢当前的房子，那么之前的房子就可以抢也可以不抢，所以将 preRob 和 preNotRob 中的较大值赋给 notRob
**/
func rob3(nums []int) int {
	var rob, notRob, preRob, preNotRob int
	for i := 0; i <= len(nums)-1; i++ {
		preNotRob, preRob = notRob, rob
		rob = preNotRob + nums[i]
		notRob = max(preNotRob, preRob)
	}

	return max(rob, notRob)
}
