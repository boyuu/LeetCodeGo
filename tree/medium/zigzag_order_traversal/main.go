package main

import (
	"container/list"
	"fmt"

	"github.com/boyuu/LeetCodeGo/tree/utils"
)

func main() {
	root := utils.MakeTreeNode(3, 9, 20, nil, nil, 15, 7)
	fmt.Println(zigzagLevelOrder(root))
}

type TreeNode = utils.TreeNode

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)

	reverse := false
	l := list.New()
	l.PushBack(root)
	for l.Len() != 0 {
		length := l.Len()
		tmp := make([]int, 0, length)
		for i := 0; i < length; i++ {
			switch reverse {
			case false:
				node := l.Remove(l.Front()).(*TreeNode)
				tmp = append(tmp, node.Val)
				if node.Left != nil {
					l.PushBack(node.Left)
				}
				if node.Right != nil {
					l.PushBack(node.Right)
				}
			case true:
				node := l.Remove(l.Back()).(*TreeNode)
				tmp = append(tmp, node.Val)
				if node.Right != nil {
					l.PushFront(node.Right)
				}
				if node.Left != nil {
					l.PushFront(node.Left)
				}
			}
		}
		result = append(result, tmp)
		switch reverse {
		case true:
			reverse = false
		case false:
			reverse = true
		}
	}
	return result
}
