package dp

import (
	"fmt"
	"testing"
)

func TestClimbStairs(t *testing.T) {

	result := climbStair(5)
	fmt.Println(result)
	result = climbStair2(5)
	fmt.Println(result)
	result = climbStair3(44)
	fmt.Println(result)
	result = climbStair4(44)
	fmt.Println(result)
}
