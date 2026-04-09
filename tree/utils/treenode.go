package utils

import (
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PtrInt(i int) *int {
	return &i
}

func MakeTreeNode(nums ...*int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := &TreeNode{Val: *nums[0]}
	stack := list.New()
	stack.PushBack(root)
	i := 1
	for stack.Len() > 0 {
		node := stack.Remove(stack.Front()).(*TreeNode)
		if i < len(nums) {
			if nums[i] != nil {
				node.Left = &TreeNode{Val: *nums[i]}
				stack.PushBack(node.Left)
			}
			i++
		}
		if i < len(nums) {
			if nums[i] != nil {
				node.Right = &TreeNode{Val: *nums[i]}
				stack.PushBack(node.Right)
			}
			i++
		}
	}
	return root
}
