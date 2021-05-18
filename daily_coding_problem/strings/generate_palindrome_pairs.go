/*
Generate palindrome pairs
Give a list ['code', 'edoc', 'da', 'd'] return [(0,1),(1,0),(2,3)]
*/
package main

import "fmt"

type (
	Pair struct {
		a, b int
	}
)

func reverse(word string) string {
	wordRune := []rune(word)
	j := len(wordRune) - 1
	for i := 0; i < len(wordRune)/2; i++ {
		wordRune[i], wordRune[j] = wordRune[j], wordRune[i]
		j--
	}
	return string(wordRune)
}

func isPalindrome(word string) bool {
	if word == reverse(word) {
		return true
	}
	return false
}

func main() {
	mp := make(map[string]int)
	words := []string{"code", "edoc", "da", "d"}

	//storing the index of each word in map
	for i := 0; i < len(words); i++ {
		mp[words[i]] = i
	}
	fmt.Println(mp)
	result := []Pair{}
	for i, word := range words {
		for j := 0; j < len(word); j++ {

			prefix, suffix := word[:j], word[j:]
			reversePrefix, reverseSuffix := reverse(prefix), reverse(suffix)
			if isPalindrome(suffix) {
				if _, ok := mp[reversePrefix]; ok {
					if i != mp[reversePrefix] {
						result = append(result, Pair{i, mp[reversePrefix]})
					}
				}
			}

			if isPalindrome(prefix) {
				if _, ok := mp[reverseSuffix]; ok {
					if i != mp[reverseSuffix] {
						result = append(result, Pair{mp[reverseSuffix], i})
					}
				}
			}

		}
	}
	fmt.Println(result)
}
