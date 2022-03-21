package main

import "fmt"

var mp = map[int]string{
	1: "", 2: "abc", 3: "def", 4: "ghi", 5: "jkl", 6: "mno", 7: "pqrs", 8: "tuv", 9: "wxyz",
}

func letter_combination_of_phone_number(str string) []string {
	if str == "" {
		return nil
	}
	res := []string{}
	backtrack(str, &res, 0, "")
	return res
}

func backtrack(str string, res *[]string, start int, t string) {
	if start == len(str) {
		*res = append(*res, append([]string{}, t)...)
		return
	}

	chars := mp[int(str[start]-'0')]
	for i := 0; i < len(chars); i++ {
		temp := t + string(chars[i])
		backtrack(str, res, start+1, temp)
		temp = temp[:len(temp)-1]
	}
}

func main() {
	fmt.Println(letter_combination_of_phone_number("23"))
}
