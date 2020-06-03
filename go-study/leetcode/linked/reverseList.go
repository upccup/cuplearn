package linked

/**
@leetcode: https://leetcode.com/problems/reverse-linked-list/
**/
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
