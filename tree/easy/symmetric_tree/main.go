package main

import (
	"fmt"

	"github.com/boyuu/LeetCodeGo/tree/utils"
)

func main() {
	root := utils.MakeTreeNode(1, 2, 2, 3, 4, 4, 3)
	fmt.Println(isSymmetric(root))
}

type TreeNode = utils.TreeNode

func isSymmetric(root *TreeNode) bool {
	return iterative(root)
}

func recursive(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return recursive0(root.Left, root.Right)
}

func recursive0(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return recursive0(left.Left, right.Right) && recursive0(left.Right, right.Left)
}

func iterative(root *TreeNode) bool {
	if root == nil {
		return true
	}

	queue := []*TreeNode{root.Left, root.Right}

	for len(queue) != 0 {
		node1 := queue[0]
		node2 := queue[1]
		queue = queue[2:]
		if node1 == nil && node2 == nil {
			continue
		}
		if node1 == nil || node2 == nil || node1.Val != node2.Val {
			return false
		}
		queue = append(queue, node1.Left)
		queue = append(queue, node2.Right)
		queue = append(queue, node1.Right)
		queue = append(queue, node2.Left)
	}
	return true
}
