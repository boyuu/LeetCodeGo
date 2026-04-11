package main

import (
	"fmt"
	"math"

	"github.com/boyuu/LeetCodeGo/tree/utils"
)

func main() {
	root := utils.MakeTreeNode(5, 1, 4, nil, nil, 3, 6)
	fmt.Println(isValidBST(root))
}

type TreeNode = utils.TreeNode

func isValidBST(root *TreeNode) bool {
	return isValidBST0(root, math.MinInt, math.MaxInt)
}

func isValidBST0(root *TreeNode, lowerBound, higherBound int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lowerBound || root.Val >= higherBound {
		return false
	}
	return isValidBST0(root.Left, lowerBound, root.Val) &&
		isValidBST0(root.Right, root.Val, higherBound)
}
