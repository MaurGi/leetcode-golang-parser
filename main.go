package main

import (
	"fmt"

	"github.com/maurgi/leetcode-golang-parser/parselist"
)

func main() {
	test := "[1,2,3]"
	res := parselist.ParseList(test)
	fmt.Println(parselist.SerializeList(res))
}
