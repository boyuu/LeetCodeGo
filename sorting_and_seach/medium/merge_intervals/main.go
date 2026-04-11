package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(merge([][]int{
		{1, 3}, {2, 6}, {8, 10}, {15, 18},
	}))
}

func merge(intervals [][]int) [][]int {
	result := make([][]int, 0)

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for _, interval := range intervals {
		if len(result) == 0 {
			result = append(result, interval)
			continue
		}
		lastInterval := result[len(result)-1]
		if interval[0] <= lastInterval[1] {
			result[len(result)-1][1] = maxInt(interval[1], lastInterval[1])
		} else {
			result = append(result, interval)
		}
	}
	return result
}

func maxInt(i, j int) int {
	if i > j {
		return i
	}
	return j
}
