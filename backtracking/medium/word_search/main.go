package main

import (
	"fmt"
)

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	fmt.Println(exist(board, "ABCCEE"))
	fmt.Println(exist(board, "SEE"))
	fmt.Println(exist(board, "ABCB"))
}

type Index struct {
	i, j int
}

func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	wordBytes := []byte(word)
	visited := make(map[Index]bool)
	for i := range board {
		for j := range board[i] {
			if board[i][j] == wordBytes[0] && dfs(board, i, j, wordBytes, visited) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, i, j int, wordBytes []byte, visited map[Index]bool) bool {
	if len(wordBytes) == 0 {
		return true
	}
	if i < 0 || i >= len(board) {
		return false
	}
	if j < 0 || j >= len(board[0]) {
		return false
	}
	idx := Index{i: i, j: j}
	if visited[idx] {
		return false
	}
	if board[i][j] != wordBytes[0] {
		return false
	}
	visited[idx] = true
	result := dfs(board, i+1, j, wordBytes[1:], visited) ||
		dfs(board, i-1, j, wordBytes[1:], visited) ||
		dfs(board, i, j+1, wordBytes[1:], visited) ||
		dfs(board, i, j-1, wordBytes[1:], visited)
	delete(visited, idx)
	return result
}
