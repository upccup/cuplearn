package linked

/**
@leetcode https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/
**/
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
