package main

import "fmt"

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
}

func containsDuplicate(nums []int) bool {
	dedup := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := dedup[num]; ok {
			return true
		}
		dedup[num] = struct{}{}
	}
	return false
}
