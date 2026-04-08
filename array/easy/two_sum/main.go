package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(nums []int, target int) []int {
	mValueIdx := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		remain := target - nums[i]
		if idx, ok := mValueIdx[remain]; ok {
			return []int{idx, i}
		}
		mValueIdx[nums[i]] = i
	}
	return nil
}
