package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)

	for _, v := range strs {
		// string to array of bytes, i.e. ['e','a','t']
		r := []rune(v)
		// sorted array of bytes, i.e. ['a','e','t']
		sort.Slice(r, func(i, j int) bool {
			return r[i] < r[j]
		})
		// sorted array of bytes back to string, i.e. "aet"
		key := string(r)
		// add the original stirng "eat" to the map with key "aet" -> ["eat", "tea"]
		m[key] = append(m[key], v)
	}

	var result [][]string
	for _, v := range m {
		result = append(result, v)
	}

	return result
}

func main() {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := groupAnagrams(input)
	fmt.Println(result)
}
