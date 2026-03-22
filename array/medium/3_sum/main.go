package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(nums)
	fmt.Println(res)
}

// 去重比较关键，下面有两个去重的地方，或者用hashmap/set去重
func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	res := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] { //去重
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			if nums[j]+nums[k] < -nums[i] {
				j++
			} else if nums[j]+nums[k] > -nums[i] {
				k--
			} else if nums[j]+nums[k] == -nums[i] {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				for j < k && nums[j] == nums[j-1] { //去重
					j++
				}
			}
		}
	}
	return res
}
