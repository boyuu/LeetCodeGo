package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}

func sortColors(nums []int) {
	idx0, idx2 := 0, len(nums)-1
	for i := 0; i <= idx2; {
		switch nums[i] {
		case 0:
			nums[i], nums[idx0] = nums[idx0], nums[i]
			idx0++
		case 1:
			i++
			continue
		case 2:
			nums[i], nums[idx2] = nums[idx2], nums[i]
			idx2--
		}
		i = idx0
	}
}
