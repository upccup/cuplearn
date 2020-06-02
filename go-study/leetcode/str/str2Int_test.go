package str

import (
	"fmt"
	"testing"
)

func TestStr2Int(t *testing.T) {
	str1 := "  dda1111222"
	num := str2Int(str1)
	fmt.Printf("str1: '%s', num: %d \n", str1, num)
	str1 = "-12211111dddd"
	num = str2Int(str1)
	fmt.Printf("str1: '%s', num: %d \n", str1, num)
	str1 = "3.141"
	num = str2Int(str1)
	fmt.Printf("str1: '%s', num: %d \n", str1, num)
}
