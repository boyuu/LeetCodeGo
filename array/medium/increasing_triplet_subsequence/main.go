package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(increasingTriplet([]int{2, 1, 5, 0, 4, 6}))
}

func increasingTriplet(nums []int) bool {
	first, second := math.MaxInt, math.MaxInt
	for i := 0; i < len(nums); i++ {
		switch {
		case nums[i] < first:
			first = nums[i]
		case nums[i] > first && nums[i] < second:
			second = nums[i]
		case nums[i] > second:
			return true
		}
	}
	return false
}
