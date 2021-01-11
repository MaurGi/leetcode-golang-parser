package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	res := 0

	for i := 0; i < len(s); i++ {
		found := make([]int, 128)
		curlen := 1
		found[int(s[i])] = 1
		for j := i + 1; j < len(s); j++ {
			if found[int(s[j])] > 0 {
				break
			} else {
				found[int(s[j])] = 1
				curlen++
			}
		}
		if curlen > res {
			res = curlen
		}
	}

	return res
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
