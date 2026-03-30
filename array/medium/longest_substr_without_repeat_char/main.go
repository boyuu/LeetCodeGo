package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

func lengthOfLongestSubstring(s string) int {
	result := 0
	mCharIdx := make(map[byte]int)

	chars := []byte(s)
	var l, r int
	for r = 0; r < len(chars); r++ {
		idx, ok := mCharIdx[chars[r]]
		if ok {
			for ; l <= idx; l++ {
				delete(mCharIdx, chars[l])
			}
		}
		mCharIdx[chars[r]] = r
		result = maxInt(result, r-l+1)
	}
	return result
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
