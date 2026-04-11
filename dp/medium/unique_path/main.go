package main

import (
	"fmt"
)

func main() {
	fmt.Println(uniquePaths(3, 7))
}

func uniquePaths(m int, n int) int {
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				dp[j] = 1
				continue
			}
			dp[j] = dp[j] + dp[j-1]
		}
	}
	return dp[n-1]
}
