package main

import "fmt"

func count_prefixes_of_a_given_string(words []string, s string) int {
	hash := make(map[string]int)
	for i := 0; i < len(words); i++ {
		hash[words[i]]++
	}
	res := 0
	for i := len(s) - 1; i >= 0; i-- {
		if hash[s[:i+1]] > 0 {
			res += hash[s[:i+1]]
		}
	}
	return res
}

func main() {
	fmt.Println(count_prefixes_of_a_given_string([]string{"a", "b", "c", "ab", "bc", "abc"}, "abc"))
}
