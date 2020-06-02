package str

import (
	"math"
	"strings"
)

/**
@leetcode https://leetcode.com/problems/string-to-integer-atoi/
**/
func str2Int(str string) int {
	str = strings.Trim(str, " ")
	isNegative := false
	var num int
	var val int

	chars := []rune(str)
	if len(chars) <= 0 {
		return 0
	}
	if chars[0] == '-' {
		isNegative = true
		chars = chars[1:]
	} else if chars[0] == '+' {
		isNegative = false
		chars = chars[1:]
	}

	for _, char := range chars {
		val = int(char - '0')
		if val < 0 || val > 9 {
			break
		}

		if 0 <= val && val <= 9 {
			num = 10*num + val
		}

		if num > math.MaxInt32 {
			if isNegative {
				num = math.MaxInt32 + 1
			} else {
				num = math.MaxInt32
			}
			break
		}
	}

	if isNegative {
		num = 0 - num
	}

	return num
}
