package dp

import (
	"fmt"
	"testing"
)

func TestRob(t *testing.T) {
	result := rob([]int{2, 7, 9, 3, 1})
	fmt.Println(result)
	result = rob2([]int{2, 7, 9, 3, 1})
	fmt.Println(result)
	result = rob3([]int{2, 7, 9, 3, 1})
	fmt.Println(result)
}
