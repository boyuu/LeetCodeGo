package main

import (
	"fmt"
)

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	i, j := 0, len(nums)-1
	for i <= j {
		mid := i + (j-i)/2
		switch {
		case nums[mid] == target:
			return mid
		case nums[i] <= nums[mid]:
			if target <= nums[mid] && target >= nums[i] {
				j = mid - 1
			} else {
				i = mid + 1
			}
		case nums[mid] <= nums[j]:
			if target >= nums[mid] && target <= nums[j] {
				i = mid + 1
			} else {
				j = mid - 1
			}
		}
	}
	return -1
}
