package main

import (
	"fmt"
	"strings"
)

func reverse_words_in_a_string(s string) string {
	str := strings.Split(s, " ")
	i, k := 0, 0
	for i < len(str) {
		if str[i] != "" {
			str[k] = str[i]
			k++
		}
		i++
	}
	str = str[:k]
	for i := 0; i < k/2; i++ {
		str[i], str[k-1-i] = str[k-1-i], str[i]
	}
	return strings.Join(str, " ")
}

func main() {
	fmt.Println(reverse_words_in_a_string("the sky is blue"))
}
