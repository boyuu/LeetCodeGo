package main

import (
	"fmt"
)

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
}

func canJump(nums []int) bool {
	maxReachable := 0
	for i, num := range nums {
		if i > maxReachable {
			return false
		}
		maxReachable = maxInt(maxReachable, i+num)
		if maxReachable >= len(nums)-1 {
			return true
		}
	}
	return true
}

func maxInt(i, j int) int {
	if i > j {
		return i
	}
	return j
}
