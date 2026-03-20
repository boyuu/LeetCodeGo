package main

import "fmt"

func main() {
	nums := []int{1, 1, 2}
	count := removeDuplicates(nums)
	for i := 0; i < count; i++ {
		fmt.Printf("%d ", nums[i])
	}
	fmt.Println("")
}
func removeDuplicates(nums []int) int {
	var curr int
	for i := 0; i < len(nums); i++ {
		if curr > 0 && nums[curr-1] == nums[i] {
			continue
		}
		nums[curr] = nums[i]
		curr++
	}
	return curr
}
