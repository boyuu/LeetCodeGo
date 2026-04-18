package main

import (
	"fmt"
)

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

func subsets(nums []int) [][]int {
	result := [][]int{{}}
	for _, num := range nums {
		result0 := copyResult(result)
		for i := range result0 {
			result0[i] = append(result0[i], num)
		}
		result = append(result, result0...)
	}
	return result
}

func copyResult(result [][]int) [][]int {
	result0 := make([][]int, 0, len(result))
	for _, rst := range result {
		result0 = append(result0, copySlice(rst))
	}
	return result0
}

func copySlice(s []int) []int {
	if s == nil {
		return nil
	}
	s0 := make([]int, len(s))
	copy(s0, s)
	return s0
}
