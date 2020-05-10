package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
@leetcode: https://leetcode.com/problems/reverse-linked-list/
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

	reverseNode := reverseListLinked(&node1)

	for reverseNode != nil {
		fmt.Println(reverseNode.Val)
		reverseNode = reverseNode.Next
	}

	testReverseListLinkedRecursion()
}

func testReverseListLinkedRecursion() {
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

func reverseListLinked(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	cur := head.Next
	pre := head

	for cur != nil {
		pre.Next = cur.Next
		cur.Next = head
		head = cur
		cur = pre.Next
	}

	return head
}

func reverseListLinkedRecursion(head, originalHead *ListNode) (*ListNode, *ListNode) {
	if originalHead == nil || originalHead.Next == nil {
		return head, originalHead
	}

	next := originalHead.Next
	originalHead.Next = next.Next
	next.Next = head
	return reverseListLinkedRecursion(next, originalHead)
}
