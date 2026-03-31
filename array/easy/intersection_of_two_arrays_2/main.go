package main

import "fmt"

func main() {
	fmt.Println(intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums2) < len(nums1) {
		nums1, nums2 = nums2, nums1
	}
	result := make([]int, 0)
	intCount := make(map[int]int)
	for _, num := range nums1 {
		intCount[num]++
	}
	for _, num := range nums2 {
		if count := intCount[num]; count > 0 {
			result = append(result, num)
			intCount[num]--
		}
	}
	return result
}
