package str

/**
https://leetcode.com/problems/longest-palindromic-substring/
**/
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}

	var beigin, length int
	length = 1
	for i := 0; i < len(s); i++ {
		// 回文串为偶数 abba 这种
		if i+1 < len(s) && s[i] == s[i+1] {
			for j := 1; i+j < len(s); j++ {
				if i-j+1 >= 0 && s[i-j+1] == s[i+j] {
					if 2*j > length {
						beigin = i - j + 1
						length = 2 * j
					}
				} else {
					break
				}
			}
		}

		// 回文串为奇数 aba 这种 奇数偶数情况可能共存所以不能用else
		if i >= 1 && i+1 < len(s) && s[i-1] == s[i+1] {
			for j := 1; i+j < len(s); j++ {
				if i-j >= 0 && s[i-j] == s[i+j] {
					if 2*j+1 > length {
						beigin = i - j
						length = 2*j + 1
					}
				} else {
					break
				}
			}

		}
	}

	return s[beigin : beigin+length]
}

/**
马拉车算法 (速度居然比上面慢...)
https://www.jianshu.com/p/392172762e55
**/
func manacher(s string) string {
	length := len(s)
	if length == 0 {
		return s
	}

	t := "$#"
	for _, v := range s {
		t += string(v)
		t += "#"
	}

	p := make([]int, len(t))
	var id, mx, resId, resMx int
	for i := 0; i < len(t); i++ {
		if mx > i {
			if p[2*id-i] < mx-i {
				p[i] = p[2*id-i]
			} else {
				p[i] = mx - i
			}
		} else {
			p[i] = 1
		}

		for i >= p[i] && i+p[i] < len(t) && t[i+p[i]] == t[i-p[i]] {
			p[i] = p[i] + 1
		}

		if mx < i+p[i] {
			mx = i + p[i]
			id = i
		}

		if resMx < p[i] {
			resMx = p[i]
			resId = i
		}
	}

	start := (resId - resMx) / 2
	return s[start : start+resMx-1]
}
