package main

import (
	"container/list"
	"fmt"

	"github.com/boyuu/LeetCodeGo/tree/utils"
)

func main() {
	root := utils.MakeTreeNode(3, 9, 20, nil, nil, 15, 7)
	fmt.Println(levelOrder(root))
}

type TreeNode = utils.TreeNode

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}

	queue := list.New()
	queue.PushBack(root)
	for queue.Len() != 0 {
		row := make([]int, 0)
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			row = append(row, node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		result = append(result, row)
	}
	return result
}
