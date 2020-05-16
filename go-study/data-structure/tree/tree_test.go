package tree

import (
	"testing"
)

func TestMidTraverse(t *testing.T) {
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 5}
	node6 := &TreeNode{Val: 6}
	node7 := &TreeNode{Val: 7}
	node8 := &TreeNode{Val: 8}
	node9 := &TreeNode{Val: 9}

	node1.LeftChild = node2
	node1.RightChild = node3
	node2.LeftChild = node4
	node2.RightChild = node5
	node3.LeftChild = node6
	node3.RightChild = node7
	node7.RightChild = node8
	node6.LeftChild = node9

	midTraverse(node1)
	midTraverseRecursion(node1)
}
