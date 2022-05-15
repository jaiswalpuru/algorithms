package main

import "fmt"
import "strings"
import "sort"

func remove_anagrams(words []string) []string {
	stack := []string{words[0]}
	i := 1
	for len(stack) > 0 && i < len(words) {
		word := strings.Split(words[i], "")
		sort.Strings(word)
		w := strings.Join(word, "")
		word2 := strings.Split(stack[len(stack)-1], "")
		sort.Strings(word2)
		w2 := strings.Join(word2, "")
		if w2 != w {
			stack = append(stack, words[i])
		}
		i++
	}
	return stack
}

func main() {
	fmt.Println(remove_anagrams([]string{
		"abba", "baba", "bbaa", "cd", "cd",
	}))
}
