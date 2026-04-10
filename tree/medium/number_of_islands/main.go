package main

import (
	"fmt"
)

// this on leetcode may have bug
func main() {
	fmt.Println(numIslands([][]byte{
		{1, 0, 1, 1, 1},
		{0, 0, 0, 0, 1},
		{0, 0, 1, 0, 1},
		{0, 0, 1, 1, 1},
	}))
}

func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])

	var result int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] != 0 {
				fillZero(grid, i, j, m, n)
				result++
			}
		}
	}
	return result
}

func fillZero(grid [][]byte, i, j, m, n int) {
	if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
		return
	}
	grid[i][j] = 0
	fillZero(grid, i-1, j, m, n)
	fillZero(grid, i+1, j, m, n)
	fillZero(grid, i, j-1, m, n)
	fillZero(grid, i, j+1, m, n)
}
