package main

import (
	"fmt"
	"strings"
)

func detect_capital(word string) bool {
	u, l := strings.ToUpper(word), strings.ToLower(word)
	if l == word {
		return true
	}
	if u == word {
		return true
	}
	first_letter := string(u[0])
	if first_letter == string(word[0]) {
		return l[1:] == word[1:]
	}
	return false
}

func main() {
	fmt.Println(detect_capital("USA"))
}
