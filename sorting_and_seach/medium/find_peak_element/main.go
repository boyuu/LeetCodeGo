package main

import (
	"fmt"
)

func main() {
	fmt.Println(findPeakElement([]int{1, 2, 3, 1}))
}

func findPeakElement(nums []int) int {
	i, j := 0, len(nums)-1
	for i < j {
		mid := i + (j-i)/2
		if nums[mid] > nums[mid+1] {
			j = mid
		} else if nums[mid] < nums[mid+1] {
			i = mid + 1
		}
	}
	return i
}
