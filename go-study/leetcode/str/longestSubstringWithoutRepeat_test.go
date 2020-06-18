package str

import (
	"fmt"
	"testing"
)

func TestLongestSubstringWithoutReapeat(t *testing.T) {
	result := longestSubstringWithoutRepeat("aba")
	fmt.Println(result)
	result = longestSubstringWithoutRepeat("au")
	fmt.Println(result)
}
