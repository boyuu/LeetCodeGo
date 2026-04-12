package main

import (
	"fmt"
)

func main() {
	fmt.Println(searchMatrix([][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}, 5))
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	m := len(matrix)
	if len(matrix[0]) == 0 {
		return false
	}
	n := len(matrix[0])

	i, j := 0, n-1
	for i >= 0 && i < m && j >= 0 && j < n {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			j--
		} else if matrix[i][j] < target {
			i++
		}
	}
	return false
}
