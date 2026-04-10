package main

import (
	"container/list"
	"fmt"

	"github.com/boyuu/LeetCodeGo/tree/utils"
)

func main() {
	root := utils.MakeTreeNode(3, 1, 4, nil, 2)
	fmt.Println(kthSmallest(root, 1))
}

type TreeNode = utils.TreeNode

func kthSmallest(root *TreeNode, k int) int {
	p := root
	stack := list.New()

	for p != nil || stack.Len() != 0 {
		for p != nil {
			stack.PushFront(p)
			p = p.Left
		}
		node := stack.Remove(stack.Front()).(*TreeNode)
		k--
		if k == 0 {
			return node.Val
		}
		p = node.Right
	}
	return -1
}
