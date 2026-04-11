package main

import (
	"fmt"
)

func main() {
	ints := searchRange([]int{2, 2}, 1)
	fmt.Println(ints)
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	return []int{binarySearchLeft(nums, 0, len(nums)-1, target),
		binarySearchRight(nums, 0, len(nums)-1, target)}
}

func binarySearchLeft(nums []int, start, end, target int) int {
	if end < start {
		return -1
	}
	if start == end {
		if nums[start] == target {
			return start
		} else {
			return -1
		}
	}
	mid := start + (end-start)/2
	switch {
	case nums[mid] >= target:
		if nums[mid] == target && (mid == 0 || nums[mid-1] < target) {
			return mid
		}
		return binarySearchLeft(nums, start, mid-1, target)
	case nums[mid] < target:
		return binarySearchLeft(nums, mid+1, end, target)
	}
	return -1
}

func binarySearchRight(nums []int, start, end, target int) int {
	if end < start {
		return -1
	}
	if start == end {
		if nums[start] == target {
			return start
		} else {
			return -1
		}
	}
	mid := start + (end-start)/2
	switch {
	case nums[mid] <= target:
		if nums[mid] == target && (mid == len(nums)-1 || nums[mid+1] > target) {
			return mid
		}
		return binarySearchRight(nums, mid+1, end, target)
	case nums[mid] > target:
		return binarySearchRight(nums, start, mid-1, target)
	}
	return -1
}
