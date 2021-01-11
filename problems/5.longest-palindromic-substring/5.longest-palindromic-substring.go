package main

import "fmt"

func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		lpali := i
		rpali := i + 1
		for ; rpali < len(s); rpali++ {
			if s[lpali] != s[rpali] {
				break
			} else {
				i++
			}
		}
		pali := s[lpali:rpali]
		for l, r := lpali-1, rpali; l >= 0 && r < len(s); l, r = l-1, r+1 {
			if s[l] == s[r] {
				pali = s[l : r+1]
			} else {
				break
			}
		}
		if len(pali) > len(res) {
			res = pali
		}
	}
	return res
}

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("a"))
	fmt.Println(longestPalindrome("ac"))
	fmt.Println(longestPalindrome("tattarrattat"))
}
