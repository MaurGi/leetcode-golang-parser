package parselist

import (
	"encoding/json"
	"log"
	"strconv"
)

// ListNode exported list node type
type ListNode struct {
	Val  int
	Next *ListNode
}

// ParseList parses a string like [1,2,3] and returns a ListNode
func ParseList(input string) *ListNode {
	var head *ListNode = nil
	var tail *ListNode = nil

	var ints []int
	err := json.Unmarshal([]byte(input), &ints)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range ints {
		newNode := &ListNode{v, nil}
		if head == nil {
			head = newNode
		} else {
			tail.Next = newNode
		}
		tail = newNode
	}

	return head
}

// SerializeList prints output
func SerializeList(input *ListNode) string {
	res := ""
	cur := input
	res += "["
	for input != nil {
		if cur != input {
			res += ","
		}
		res += strconv.Itoa(input.Val)
		input = input.Next
	}
	res += "]"

	return res
}
