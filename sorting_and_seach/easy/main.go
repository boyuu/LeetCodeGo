package main

import (
	"fmt"
)

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	idx := m + n - 1
	i, j := m-1, n-1

	for i >= 0 || j >= 0 {
		if i < 0 {
			nums1[idx] = nums2[j]
			idx--
			j--
			continue
		}
		if j < 0 {
			nums1[idx] = nums1[i]
			idx--
			i--
			continue
		}
		if nums1[i] >= nums2[j] {
			nums1[idx] = nums1[i]
			idx--
			i--
		} else {
			nums1[idx] = nums2[j]
			idx--
			j--
		}
	}
}
