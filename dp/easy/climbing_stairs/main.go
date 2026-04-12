package main

import (
	"fmt"
)

func main() {
	fmt.Println(climbStairs(5))
}

func climbStairs(n int) int {
	prev2, prev1 := 0, 1
	var result int
	for i := 0; i < n; i++ {
		result = prev2 + prev1
		prev2, prev1 = prev1, result
	}
	return result
}
