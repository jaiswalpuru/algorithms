package main

import "fmt"

func letter_case_permutation(s string) []string {
	res := []string{}
	backtrack(s, []byte{}, 0, &res)
	return res
}

func backtrack(str string, set []byte, start int, res *[]string) {
	if start == len(str) {
		*res = append(*res, string(set))
		return
	}
	backtrack(str, append(set, str[start]), start+1, res)
	if str[start] >= 'a' && str[start] <= 'z' {
		backtrack(str, append(set, str[start]-32), start+1, res)
	} else if str[start] >= 'A' && str[start] <= 'z' {
		backtrack(str, append(set, str[start]+32), start+1, res)
	}
}

func main() {
	fmt.Println(letter_case_permutation("a1b2"))
}
