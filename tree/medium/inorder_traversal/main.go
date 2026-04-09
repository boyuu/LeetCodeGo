package main

import (
	"container/list"
	"fmt"

	"github.com/boyuu/LeetCodeGo/tree/utils"
)

func main() {
	root := utils.MakeTreeNode(1, nil, 2, 3)
	nums0 := inorderTraversal(root)
	nums1 := inorderRecursive(root)
	fmt.Println(nums0)
	fmt.Println(nums1)
}

type TreeNode = utils.TreeNode

func inorderTraversal(root *TreeNode) []int {
	return inorderIterative(root)
}

func inorderRecursive(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	nums := inorderRecursive(root.Left)
	nums = append(nums, root.Val)
	nums = append(nums, inorderRecursive(root.Right)...)
	return nums
}

func inorderIterative(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)

	stack := list.New()
	p := root
	for p != nil || stack.Len() != 0 {
		for p != nil {
			stack.PushFront(p)
			p = p.Left
		}
		node := stack.Remove(stack.Front()).(*TreeNode)
		result = append(result, node.Val)
		p = node.Right
	}
	return result
}
