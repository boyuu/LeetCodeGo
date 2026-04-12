package main

import (
	"fmt"
)

func main() {
	fmt.Println(coinChange([]int{1, 5, 10}, 11))
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			remain := i - coin
			if remain >= 0 {
				dp[i] = minInt(dp[i], dp[remain]+1)
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func minInt(i, j int) int {
	if i < j {
		return i
	}
	return j
}
