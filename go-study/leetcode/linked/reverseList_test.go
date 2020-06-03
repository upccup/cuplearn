package linked

import (
	"fmt"
	"testing"
)

func TestReveseList(t *testing.T) {
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

	reverseNode := reverseListLinked(&node1)

	for reverseNode != nil {
		fmt.Println(reverseNode.Val)
		reverseNode = reverseNode.Next
	}
}

func TestReverseListLinkedRecursion(t *testing.T) {
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

	fmt.Println("reverseListLinkedRecursion")
	reverseNode2, _ := reverseListLinkedRecursion(&node1, &node1)
	for reverseNode2 != nil {
		fmt.Println(reverseNode2.Val)
		reverseNode2 = reverseNode2.Next
	}
}
