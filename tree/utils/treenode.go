package utils

import (
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MakeTreeNode(nums ...any) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	for _, num := range nums {
		if num == nil {
			continue
		}
		if _, ok := num.(int); !ok {
			panic("nums element should be either int or nil")
		}
	}
	root := &TreeNode{Val: nums[0].(int)}
	stack := list.New()
	stack.PushBack(root)
	i := 1
	for stack.Len() > 0 {
		node := stack.Remove(stack.Front()).(*TreeNode)
		if i < len(nums) {
			if nums[i] != nil {
				node.Left = &TreeNode{Val: nums[i].(int)}
				stack.PushBack(node.Left)
			}
			i++
		}
		if i < len(nums) {
			if nums[i] != nil {
				node.Right = &TreeNode{Val: nums[i].(int)}
				stack.PushBack(node.Right)
			}
			i++
		}
	}
	return root
}
