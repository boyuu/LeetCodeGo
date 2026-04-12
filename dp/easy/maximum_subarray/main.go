package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxSum := nums[0]
	result := maxSum
	for i := 1; i < len(nums); i++ {
		maxSum = maxInt(maxSum+nums[i], nums[i])
		result = maxInt(result, maxSum)
	}
	return result
}

func maxInt(i, j int) int {
	if i > j {
		return i
	}
	return j
}
