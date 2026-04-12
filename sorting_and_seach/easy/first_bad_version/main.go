package main

import (
	"fmt"
)

func main() {
	fmt.Println(firstBadVersion(5))
}

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 *
 */

func isBadVersion(version int) bool {
	if version >= 4 {
		return true
	}
	return false
}

func firstBadVersion(n int) int {
	i, j := 1, n
	for i <= j {
		mid := i + (j-i)/2
		switch isBadVersion(mid) {
		case true:
			if mid == 1 || !isBadVersion(mid-1) {
				return mid
			}
			j = mid - 1
		case false:
			i = mid + 1
		}
	}
	return -1
}
