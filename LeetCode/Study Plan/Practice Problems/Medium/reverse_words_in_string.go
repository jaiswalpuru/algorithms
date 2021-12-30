/*
Given an input string s, reverse the order of the words.

A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.

Return a string of the words in reverse order concatenated by a single space.

Note that s may contain leading or trailing spaces or multiple spaces between two words. The returned string
should only have a single space separating the words. Do not include any extra spaces.
*/

package main

import "fmt"

func reverse_words(str string) string {
	res := ""
	n := len(str)
	words := make([]string, 0)
	temp := ""
	for i := 0; i < n; i++ {
		if str[i] >= 'a' && str[i] <= 'z' || str[i] >= 'A' && str[i] <= 'Z' || str[i] >= '0' && str[i] <= '9' {
			temp += string(str[i])
		} else {
			if temp != "" {
				words = append(words, temp)
			}
			temp = ""
		}
	}
	if temp != " " && temp != "" {
		words = append(words, temp)
	}

	//reverse the words
	m := len(words)
	for i := m - 1; i >= 0; i-- {
		if i == 0 {
			res += words[i]
		} else {
			res += words[i] + " "
		}
	}

	return res
}

func main() {
	fmt.Println(reverse_words("the sky is blue"))
	fmt.Println(reverse_words("  Bob    Loves  Alice   "))
}
