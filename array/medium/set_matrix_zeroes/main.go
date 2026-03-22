package main

import "fmt"

func main() {
	//matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(matrix)
	fmt.Println(matrix)
}

func setZeroes(matrix [][]int) {
	rows := make(map[int]struct{})
	cols := make(map[int]struct{})
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				rows[i] = struct{}{}
				cols[j] = struct{}{}
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			_, ok := rows[i]
			_, ok0 := cols[j]
			if ok || ok0 {
				matrix[i][j] = 0
			}
		}
	}
}
