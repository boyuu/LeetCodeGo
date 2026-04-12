package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}

func lengthOfLIS(nums []int) int {
	// return dpSolution(nums)
	return binarySearchSolution(nums)
}

func binarySearchSolution(nums []int) int {
	tails := make([]int, 0)
	for _, num := range nums {
		if len(tails) == 0 || tails[len(tails)-1] < num {
			tails = append(tails, num)
			continue
		}
		// binary search tails to find the first element greater or equal than num
		idx := binarySearch(tails, num)
		tails[idx] = num
	}
	return len(tails)
}

func binarySearch(tails []int, target int) int {
	i, j := 0, len(tails)-1
	for i <= j {
		mid := i + (j-i)/2
		switch {
		case tails[mid] >= target:
			if mid == 0 || tails[mid-1] < target {
				return mid
			}
			j = mid - 1
		case tails[mid] < target:
			i = mid + 1
		}
	}
	return -1 //impossible
}

func dpSolution(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	result := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = maxInt(dp[i], dp[j]+1)
			}
		}
		result = maxInt(result, dp[i])
	}
	return result
}

func maxInt(i, j int) int {
	if i > j {
		return i
	}
	return j
}
