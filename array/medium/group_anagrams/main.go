package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}

func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)
	grouped := make(map[string][]string)
	for _, str := range strs {
		grouped[getKey(str)] = append(grouped[getKey(str)], str)
	}
	for _, group := range grouped {
		result = append(result, group)
	}
	return result
}

func getKey(str string) string {
	strBytes := []byte(str)
	sort.Slice(strBytes, func(i, j int) bool {
		return strBytes[i] < strBytes[j]
	})
	return string(strBytes)
}
