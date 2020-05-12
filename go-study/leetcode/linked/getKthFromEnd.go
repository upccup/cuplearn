package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
@leetcode https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/
**/
func main() {
	node1 := ListNode{1, nil}
	node2 := ListNode{2, nil}
	node3 := ListNode{3, nil}
	node4 := ListNode{4, nil}
	node5 := ListNode{5, nil}
	node6 := ListNode{6, nil}

	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5
	node5.Next = &node6

	result := getKthFromEnd(&node1, 7)

	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	leftPointer := head
	rightPointer := head

	for i := 0; i < k && rightPointer != nil; i++ {
		rightPointer = rightPointer.Next
	}

	for rightPointer != nil {
		rightPointer = rightPointer.Next
		leftPointer = leftPointer.Next
	}

	return leftPointer
}
