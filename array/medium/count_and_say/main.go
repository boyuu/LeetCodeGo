package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countAndSay(4))
}

func countAndSay(n int) string {
	s := "1"
	for i := 1; i < n; i++ {
		s = rle(s)
	}
	return s
}

func rle(s string) string {
	bs := []byte(s)
	result := make([]byte, 0)
	for i := 0; i < len(bs); {
		j := i
		for j < len(bs) && bs[j] == bs[i] {
			j++
		}
		count := j - i
		result = strconv.AppendInt(result, int64(count), 10)
		result = append(result, bs[i])
		i = j
	}
	return string(result)
}
