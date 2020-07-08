package leetcode

import "math"

func ReverseInteger(x int) int {
	res := 0
	for ; x != 0; x = x / 10 {
		res = res*10 + x%10

		if res > math.MaxInt32 || 0-res > math.MaxInt32 {
			return 0
		}

	}
	return res
}
