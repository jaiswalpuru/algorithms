/*
Given a mapping of digits to letters (as in a phone number), and a digit string,
return all possible letters the number could represent. You can assume each valid number
in the mapping is a single digit.

For example if {“2”: [“a”, “b”, “c”], 3: [“d”, “e”, “f”], …} then “23” should
return [“ad”, “ae”, “af”, “bd”, “be”, “bf”, “cd”, “ce”, “cf"].
*/

package main

import (
	"fmt"
)

func produce_combination(a, b []string) []string {
	res := []string{}
	n, m := len(a), len(b)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			res = append(res, a[i]+b[j])
		}
	}
	return res
}

func possible_combinations(num_map map[string][]string, digit string) []string {
	n := len(digit)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return num_map[digit]
	}
	res := []string{}
	res = append(res, num_map[string(digit[0])]...)
	for i := 1; i < n; i++ {
		res = produce_combination(res, num_map[string(digit[i])])
	}
	return res
}

func main() {
	num_letter_map := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
	digit := "234"
	fmt.Println(possible_combinations(num_letter_map, digit))
}
