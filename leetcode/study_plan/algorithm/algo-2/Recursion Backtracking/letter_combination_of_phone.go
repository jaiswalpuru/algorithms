/*
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.
Return the answer in any order.

A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.
*/

package main

import "fmt"

var mp = map[int]string{
	1: "", 2: "abc", 3: "def", 4: "ghi", 5: "jkl", 6: "mno", 7: "pqrs", 8: "tuv", 9: "wxyz",
}

func letter_combination(digits string) []string {

	if digits == "" {
		return nil
	}

	res := []string{}
	var temp = ""
	_letter_combination(digits, &res, &temp, 0)
	return res
}

func _letter_combination(digits string, res *[]string, temp *string, start int) {

	if start > len(digits) {
		return
	}

	if len(*temp) == len(digits) {
		*res = append(*res, *temp)
		return
	}

	list := mp[int(digits[start]-'0')]
	for i := 0; i < len(list); i++ {
		*temp += string(list[i])
		_letter_combination(digits, res, temp, start+1)
		*temp = (*temp)[:len(*temp)-1]
	}
}

func main() {
	fmt.Println(letter_combination("23"))
}
