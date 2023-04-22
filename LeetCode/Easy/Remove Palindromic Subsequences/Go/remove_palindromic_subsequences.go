package main

import "fmt"

func remove_palindromic_subsequence(str string) int {
	i, j := 0, len(str)-1
	for i < j {
		if str[i] != str[j] {
			return 2
		}
		i++
		j--
	}
	return 1
}

func main() {
	fmt.Println(remove_palindromic_subsequence("ababa"))
}
