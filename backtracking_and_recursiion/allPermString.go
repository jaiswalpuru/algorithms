package main

import (
	"fmt"
)

func swap(str []rune, i, j int) {
	str[i], str[j] = str[j], str[i]
}

func stringPerm(str []rune, i int) {
	if i == len(str) {
		fmt.Println(string(str))
	}
	for j := i; j < len(str); j++ {
		swap(str, i, j)
		fmt.Println(i, string(str))
		stringPerm(str, i+1)
		swap(str, i, j)
	}
}

func main() {
	stringPerm([]rune("ABC"), 0)
}
