package tree

import (
	"fmt"

	"github.com/cuplearn/go-study/data-structure/stack"
)

type TreeNode struct {
	Val        interface{}
	LeftChild  *TreeNode
	RightChild *TreeNode
}

func midTraverse(root *TreeNode) {
	if root == nil {
		return
	}

	s := stack.New()
	p := root

	for p != nil || s.Length() > 0 {
		for p != nil {
			s.Push(p)
			p = p.LeftChild
		}

		if s.Length() > 0 {
			p = s.Pop().(*TreeNode)
			fmt.Println(p.Val)
			p = p.RightChild
		}
	}

}

func midTraverseRecursion(root *TreeNode) {
	if root != nil {
		midTraverseRecursion(root.LeftChild)
		fmt.Println(root.Val)
		midTraverseRecursion(root.RightChild)
	}
}
