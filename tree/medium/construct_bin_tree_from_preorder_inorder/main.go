package main

import (
	"github.com/boyuu/LeetCodeGo/tree/utils"
)

func main() {
	buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
}

type TreeNode = utils.TreeNode

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: preorder[0],
	}
	rootIdx := 0
	for inorder[rootIdx] != preorder[0] {
		rootIdx++
	}
	root.Left = buildTree(preorder[1:rootIdx+1], inorder[0:rootIdx])
	root.Right = buildTree(preorder[rootIdx+1:], inorder[rootIdx+1:])
	return root
}
