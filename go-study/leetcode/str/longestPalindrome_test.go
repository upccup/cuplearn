package str

import (
	"fmt"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	result := longestPalindrome("xxs")
	fmt.Println(result)
	result = longestPalindrome("babad")
	fmt.Println(result)
	result = longestPalindrome("abacab")
	fmt.Println(result)
	result = longestPalindrome("ccc")
	fmt.Println(result)
	result = longestPalindrome("aaabaaaa")
	fmt.Println(result)
	result = longestPalindrome("babadada")
	fmt.Println(result)
}

func TestManacher(t *testing.T) {
	result := manacher("xxs")
	fmt.Println(result)
	result = manacher("babad")
	fmt.Println(result)
	result = manacher("abacab")
	fmt.Println(result)
	result = manacher("ccc")
	fmt.Println(result)
	result = manacher("aaabaaaa")
	fmt.Println(result)
	result = manacher("babadada")
	fmt.Println(result)
}
