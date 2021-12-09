/*
Given a string containing digits from 2-9 inclusive,
return all possible letter combinations that the number could represent. Return the answer in any order.

A mapping of digit to letters (just like on the telephone buttons) is given below.
Note that 1 does not map to any letters.

Example 1:
Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]

Example 2:
Input: digits = ""
Output: []

Example 3:
Input: digits = "2"
Output: ["a","b","c"]
*/

package main

import "fmt"

func _phone_letter_combi(digits string, temp *string, mp map[int]string, res *[]string, start int) {

	if start > len(digits) {
		return
	}

	if len(*temp) == len(digits) {
		*res = append(*res, *temp)
		return
	}

	num := int(digits[start] - '0')
	list := mp[num]
	for i := 0; i < len(list); i++ {
		*temp += string(list[i])
		_phone_letter_combi(digits, temp, mp, res, start+1)
		*temp = (*temp)[:len(*temp)-1]
	}
}

func phone_letter_combi(digits string) []string {
	res := []string{}
	if len(digits) == 0 {
		return res
	}
	ltrs := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	j, mp := 0, make(map[int]string)
	for i := 2; i <= 9; i++ {
		mp[i] = ltrs[j]
		j++
	}
	var temp string
	_phone_letter_combi(digits, &temp, mp, &res, 0)
	return res
}

func main() {
	fmt.Println(phone_letter_combi("23"))
	fmt.Println(phone_letter_combi(""))
	fmt.Println(phone_letter_combi("2"))
}
