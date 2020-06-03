package linked

import "fmt"

// ListNode is simply link define
type ListNode struct {
	Val  int
	Next *ListNode
}

func printLinked(node *ListNode) {
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}
