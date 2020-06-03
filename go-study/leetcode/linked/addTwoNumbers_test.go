package linked

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	node1 := ListNode{1, nil}
	node2 := ListNode{2, nil}
	node3 := ListNode{3, nil}
	node4 := ListNode{4, nil}
	node5 := ListNode{5, nil}
	node6 := ListNode{6, nil}
	node7 := ListNode{7, nil}

	node1.Next = &node2
	node2.Next = &node3

	node4.Next = &node5
	node5.Next = &node6
	node6.Next = &node7

	result := addTwoNumbers(&node1, &node4)
	printLinked(result)
}

func TestAddTwoNumbers2(t *testing.T) {
	node1 := ListNode{9, nil}
	node2 := ListNode{8, nil}
	node3 := ListNode{1, nil}
	node1.Next = &node2

	result := addTwoNumbers(&node1, &node3)
	printLinked(result)
}

func TestAddTwoNumbersOptimize(t *testing.T) {
	node1 := ListNode{1, nil}
	node2 := ListNode{2, nil}
	node3 := ListNode{3, nil}
	node4 := ListNode{4, nil}
	node5 := ListNode{5, nil}
	node6 := ListNode{6, nil}
	node7 := ListNode{7, nil}

	node1.Next = &node2
	node2.Next = &node3

	node4.Next = &node5
	node5.Next = &node6
	node6.Next = &node7

	result := addTwoNumbersOptimize(&node1, &node4)
	printLinked(result)
}
