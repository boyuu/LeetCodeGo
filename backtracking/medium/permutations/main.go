package main

import (
	"fmt"
)

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	var backtrack func(remain []int, tmp []int)
	backtrack = func(remain []int, tmp []int) {
		if len(remain) == 0 {
			tmp0 := make([]int, len(tmp))
			copy(tmp0, tmp)
			result = append(result, tmp0)
			return
		}
		for i := 0; i < len(remain); i++ {
			remain[0], remain[i] = remain[i], remain[0]
			tmp = append(tmp, remain[0])
			backtrack(remain[1:], tmp)
			remain[0], remain[i] = remain[i], remain[0]
			tmp = tmp[:len(tmp)-1]
		}
	}
	backtrack(nums, make([]int, 0, len(nums)))
	return result
}
