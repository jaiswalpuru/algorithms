package main

import "fmt"

//------------brute force-------------------------
func short_encode_of_words(words []string) int {
	set := make(map[string]struct{})
	for i := 0; i < len(words); i++ {
		set[words[i]] = struct{}{}
	}

	for i := 0; i < len(words); i++ {
		for j := 1; j < len(words[i]); j++ {
            delete(set, words[i][j:])
		}
	}

	res := 0
	for k := range set {
		res += len(k) + 1
	}
	return res
}

//------------brute force-------------------------

func main() {
	fmt.Println(short_encode_of_words([]string{"time", "me", "bell"}))
}
