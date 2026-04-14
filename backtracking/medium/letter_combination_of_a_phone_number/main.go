package main

import (
	"fmt"
)

func main() {
	fmt.Println(letterCombinations("295"))
}

var mDigitToLetters = map[byte][]byte{
	'2': []byte("abc"),
	'3': []byte("def"),
	'4': []byte("ghi"),
	'5': []byte("jkl"),
	'6': []byte("mno"),
	'7': []byte("pqrs"),
	'8': []byte("tuv"),
	'9': []byte("wxyz"),
}

func letterCombinations(digits string) []string {
	result := make([]string, 0)
	var recursive func(digits []byte, s string)
	recursive = func(digits []byte, s string) {
		if len(digits) == 0 {
			result = append(result, s)
			return
		}
		digit := digits[0]
		letters := mDigitToLetters[digit]
		for _, letter := range letters {
			recursive(digits[1:], s+string(letter))
		}
	}
	recursive([]byte(digits), "")
	return result
}
