package main

import (
	"fmt"
)

func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	result := make([]string, 0)

	var backtrack func(left, right int, s string)
	backtrack = func(left, right int, s string) {
		if left == 0 && right == 0 {
			result = append(result, s)
			return
		}
		if left > 0 {
			backtrack(left-1, right, s+"(")
		}
		if right > 0 && right > left {
			backtrack(left, right-1, s+")")
		}
	}
	backtrack(n, n, "")
	return result
}
