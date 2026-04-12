package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {
	minPrice := math.MaxInt
	maxProfit0 := 0
	for _, price := range prices {
		if price-minPrice > 0 {
			maxProfit0 = maxInt(maxProfit0, price-minPrice)
		}
		minPrice = minInt(minPrice, price)
	}
	return maxProfit0
}

func maxInt(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func minInt(i, j int) int {
	if i < j {
		return i
	}
	return j
}
