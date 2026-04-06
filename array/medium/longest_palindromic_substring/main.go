package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("abcba"))
}

func longestPalindrome(s string) string {
	bs := []byte(s)
	var start, end int
	for i := 0; i < len(bs); i++ {
		l, r := i, i
		for l >= 0 && r < len(bs) && bs[l] == bs[r] {
			if r-l+1 > end-start+1 {
				start, end = l, r
			}
			l--
			r++
		}
	}
	for i := 0; i < len(bs)-1; i++ {
		l, r := i, i+1
		for l >= 0 && r < len(bs) && bs[l] == bs[r] {
			if r-l+1 > end-start+1 {
				start, end = l, r
			}
			l--
			r++
		}
	}
	return string(bs[start : end+1])
}
