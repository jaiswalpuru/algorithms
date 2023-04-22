package main

import (
	"fmt"
	"strconv"
)

func generalized_abbreviation(word string) []string {
	res := []string{}
	recur(word, []byte{}, 0, 0, &res)
	return res
}

func recur(word string, set []byte, start, cnt int, res *[]string) {
	if start == len(word) {
		if cnt != 0 {
			set = append(set, []byte(strconv.Itoa(cnt))...)
		}
		*res = append(*res, string(set))
		return
	}
	n := len(set)
	recur(word, set, start+1, cnt+1, res)
	if cnt != 0 {
		set = append(set, []byte(strconv.Itoa(cnt))...)
	}
	set = append(set, word[start])
	recur(word, set, start+1, 0, res)
	set = set[:n]
}

func main() {
	fmt.Println(generalized_abbreviation("word"))
}
