package main

import (
	"fmt"
)

func main() {
	prices := []int{7, 6, 4, 3, 1}
	profit := maxProfit(prices)
	fmt.Println(profit)
}

func maxProfit(prices []int) int {
	totalProfit := 0
	for i := range prices {
		if i == 0 {
			continue
		}
		profit := prices[i] - prices[i-1]
		if profit > 0 {
			totalProfit += profit
		}
	}
	return totalProfit
}
