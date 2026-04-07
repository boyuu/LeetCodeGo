package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{9, 9, 9}))
}

func plusOne(digits []int) []int {
	remain := 1
	i := len(digits) - 1
	for remain > 0 {
		if i < 0 {
			newDigits := make([]int, 0, len(digits)+1)
			newDigits = append(newDigits, remain)
			newDigits = append(newDigits, digits...)
			return newDigits
		}
		newDigit := (digits[i] + remain) % 10
		remain = (digits[i] + remain) / 10
		digits[i] = newDigit
		i--
	}
	return digits
}
