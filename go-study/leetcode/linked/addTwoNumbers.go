package linked

/**
https://leetcode.com/problems/add-two-numbers/
**/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var preNode, result *ListNode
	carrySign := false
	for l1 != nil && l2 != nil {
		value := l1.Val + l2.Val
		if carrySign {
			value++
		}

		if value > 9 {
			value = value - 10
			carrySign = true
		} else {
			carrySign = false
		}
		n := ListNode{value, nil}
		if preNode != nil {
			preNode.Next = &n
		} else {
			result = &n
		}

		preNode = &n
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		if !carrySign {
			preNode.Next = l1
			break
		} else {
			value := l1.Val + 1
			if value > 9 {
				carrySign = true
				l1.Val = value - 10
				preNode.Next = l1
				preNode = l1
				l1 = l1.Next
			} else {
				carrySign = false
				l1.Val = value
				preNode.Next = l1
				break
			}
		}
	}

	for l2 != nil {
		if !carrySign {
			preNode.Next = l2
			break
		} else {
			value := l2.Val + 1
			if value > 9 {
				carrySign = true
				l2.Val = value - 10
				preNode.Next = l2
				preNode = l2
				l2 = l2.Next
			} else {
				carrySign = false
				l2.Val = value
				preNode.Next = l2
				break
			}
		}
	}

	// 如果两个linked长度相等最后需要进一位
	if carrySign {
		preNode.Next = &ListNode{1, nil}
	}
	return result
}

func addTwoNumbersOptimize(l1, l2 *ListNode) *ListNode {
	beginNode := &ListNode{-1, nil}
	carrySign := 0
	curNode := beginNode

	for l1 != nil || l2 != nil {
		var val1, val2 int
		if l1 == nil {
			val1 = 0
		} else {
			val1 = l1.Val
		}

		if l2 == nil {
			val2 = 0
		} else {
			val2 = l2.Val
		}

		sum := val1 + val2 + carrySign
		if sum >= 10 {
			carrySign = 1
			sum = sum % 10
		} else {
			carrySign = 0
		}

		curNode.Next = &ListNode{sum, nil}
		curNode = curNode.Next

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carrySign == 1 {
		curNode.Next = &ListNode{1, nil}
	}

	return beginNode.Next
}
