package linked

func reverseBetween(m, n int, head *ListNode) *ListNode {
	i := 1
	var preCutNode, reverseHead, cur *ListNode

	next := head
	for i <= n {
		switch {
		case i < m-1:
			next = next.Next
			i++
		case i == m-1:
			preCutNode = next
			next = next.Next
			i++
		case i == m:
			reverseHead = next
			cur = reverseHead
			next = next.Next
			i++
		default:
			reverseHead.Next = next.Next
			next.Next = cur
			cur = next
			next = reverseHead.Next
			i++
		}
	}

	if preCutNode != nil {
		preCutNode.Next = cur
	}

	if m == 1 {
		return cur
	}

	return head
}
