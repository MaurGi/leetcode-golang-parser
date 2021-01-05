package main

import (
	"fmt"

	"github.com/maurgi/leetcode-golang-parser/parselist"
)

// problems from https://medium.com/@koheiarai94/60-leetcode-questions-to-prepare-for-coding-interview-8abbb6af589e
// https://leetcode.com/problems/reverse-linked-list/

// tail: 1 -> 2 -> 3 -> 4 -> nil
// result: nil

// tail: 2 -> 3 -> 4 -> nil
// result : 1 -> nil

// 3 -> 4
// 2 -> 1

// 4
// 3 -> 2 -> 1

func reverseList(head *parselist.ListNode) *parselist.ListNode {
	remaining := head
	var result *parselist.ListNode

	for remaining != nil {
		resulttail := result
		resulthead := remaining
		remaining = remaining.Next
		resulthead.Next = resulttail
		result = resulthead
	}

	return result
}

func main() {
	input := "[1,2,3]"
	inputList := parselist.ParseList(input)
	fmt.Println(parselist.SerializeList(inputList))

	resultList := reverseList(inputList)
	fmt.Println(parselist.SerializeList(resultList))
}
