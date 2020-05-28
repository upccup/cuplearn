package dp

import "math"

// https://leetcode.com/problems/climbing-stairs/
/**
参考文档: https://github.com/grandyang/leetcode/issues/70
爬到第n层的方式只有两种方式---从n-2层跨2步和n-1层跨一步
所以爬到第n层的路径集合就是在n-2层结果集后都加一个2,或n-1的结果集中都加一个1,由于加的数不一样这两个结果集之间就不会重复
所以到dp[n]=dp[n-1]+dp[n-2]
**/
func climbStair(n int) int {
	if n <= 1 {
		return 1
	}

	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 0, 1, 2

	for i := 2; i < n; i++ {
		dp[i+1] = dp[i] + dp[i-1]
	}

	return dp[n]
}

/**
 对空间进行进一步优化，只用两个整型变量a和b来存储过程值，
 首先将 a+b 的值赋给b，然后a赋值为原来的b，所以应该赋值为 b-a 即可。
 这样就模拟了上面累加的过程，而不用存储所有的值
**/
func climbStair2(n int) int {
	a, b := 1, 1
	for i := n; i > 0; i-- {
		b = a + b
		a = b - a
	}

	return a
}

/**
使用递归加记忆数组的方式, 因为记忆数组可以保存已经计算过的结果这样就不会存在重复计算了,大大提高了运行效率
**/
func climbStair3(n int) int {
	memo := make([]int, n+1)
	return helper(n, memo)
}

func helper(n int, memo []int) int {
	if n <= 1 {
		return 1
	}

	if memo[n] > 0 {
		return memo[n]
	}

	memo[n] = helper(n-1, memo) + helper(n-2, memo)
	return memo[n]
}

/**
使用数列的通项公式直接求值 斐波那契数列的通项公式介绍  https://zhuanlan.zhihu.com/p/26679684
**/
func climbStair4(n int) int {
	root5 := math.Sqrt(5)
	return int((1 / root5) * (math.Pow((1+root5)/2, float64(n+1)) - math.Pow((1-root5)/2, float64(n+1))))
}
