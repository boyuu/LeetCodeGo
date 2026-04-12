package main

import (
	"fmt"
)

func main() {
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	prevRob1 := nums[0]
	prevRob2 := maxInt(nums[0], nums[1])
	currRob := maxInt(prevRob1, prevRob2)
	for i := 2; i < len(nums); i++ {
		currRob = maxInt(prevRob1+nums[i], prevRob2)
		prevRob1, prevRob2 = prevRob2, currRob
	}
	return currRob
}

func maxInt(i, j int) int {
	if i > j {
		return i
	}
	return j
}
