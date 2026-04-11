package main

import (
	"fmt"

	"github.com/boyuu/LeetCodeGo/tree/utils"
)

func main() {
	root := utils.MakeTreeNode(3, 9, 20, nil, nil, 15, 7)
	fmt.Println(maxDepth(root))
}

type TreeNode = utils.TreeNode

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftMax := maxDepth(root.Left)
	rightMax := maxDepth(root.Right)
	if leftMax > rightMax {
		return leftMax + 1
	}
	return rightMax + 1
}
