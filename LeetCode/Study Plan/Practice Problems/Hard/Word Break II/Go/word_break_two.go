package main

import (
	"fmt"
	"strings"
)

func word_break_two(s string, word_dict []string) []string {
	mp := make(map[string]bool)
	for i := 0; i < len(word_dict); i++ {
		mp[word_dict[i]] = true
	}
	res := []string{}
	backtrack(s, &res, []string{}, 0, mp)
	return res
}

func backtrack(s string, res *[]string, set []string, ind int, mp map[string]bool) {
	if ind == len(s) {
		*res = append(*res, strings.Join(set, " "))
		return
	}
	t := ""
	temp := []string{}
	for i := ind; i < len(s); i++ {
		t += string(s[i])
		if mp[t] {
			temp = append(set, t)
			backtrack(s, res, temp, i+1, mp)
			temp = temp[:len(temp)-1]
		}
	}
}

func main() {
	fmt.Println(word_break_two("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
}
