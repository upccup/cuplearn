package str

/**
https://leetcode.com/problems/longest-substring-without-repeating-characters/
**/
func longestSubstringWithoutRepeat(s string) int {
	left, result, length := -1, 0, len(s)
	m := make(map[byte]int)
	var v byte

	for i := 0; i < length; i++ {
		v = s[i]
		sign, ok := m[v]
		if ok && sign > left {
			left = sign
		}

		m[v] = i
		if result < i-left {
			result = i - left
		}

	}

	return result
}
