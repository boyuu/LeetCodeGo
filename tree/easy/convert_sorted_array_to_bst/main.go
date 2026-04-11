package main

import (
	"fmt"

	"github.com/boyuu/LeetCodeGo/tree/utils"
)

func main() {
	root := sortedArrayToBST([]int{-10, -3, 0, 5, 9})
	fmt.Println(root)
}

type TreeNode = utils.TreeNode

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	rootIdx := len(nums) / 2
	left := sortedArrayToBST(nums[:rootIdx])
	right := sortedArrayToBST(nums[rootIdx+1:])
	return &TreeNode{
		Val:   nums[rootIdx],
		Left:  left,
		Right: right,
	}
}
