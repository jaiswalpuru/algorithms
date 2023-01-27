package main

import "fmt"

var numberLetterMap = map[byte]string{
	'1': "", '2': "abc", '3': "def",
	'4': "ghi", '5': "jkl", '6': "mno",
	'7': "pqrs", '8': "tuv", '9': "wxyz",
}

func letterCombinations(digits string) []string {
	res := []string{}
	backtrack(digits, 0, "", &res)
	return res
}

func backtrack(digits string, ind int, temp string, res *[]string) {
	if ind == len(digits) {
		if temp != "" {
			*res = append(*res, temp)
		}
		return
	}
	letters := numberLetterMap[digits[ind]]
	for i := 0; i < len(letters); i++ {
		temp += string(letters[i])
		backtrack(digits, ind+1, temp, res)
		temp = temp[:len(temp)-1]
	}
}

func main() {
	fmt.Println(letterCombinations("23"))
}
