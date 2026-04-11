package main

import (
	"fmt"
)

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}

func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, len(nums)-k, 0, len(nums)-1)
}

func quickSelect(nums []int, k, start, end int) int {
	pivot := nums[start]
	i, j := start+1, end
	for i <= j {
		for i <= j && nums[i] <= pivot {
			i++
		}
		for i <= j && nums[j] >= pivot {
			j--
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[start], nums[j] = nums[j], nums[start]
	switch {
	case j == k:
		return nums[j]
	case j < k:
		return quickSelect(nums, k, j+1, end)
	case j > k:
		return quickSelect(nums, k, start, j-1)
	}
	return 0
}
