package main

import "fmt"

//------------------brute force---------------------
func vowels_of_all_substrings(word string) int64 {
	res := int64(0)
	for i := 0; i < len(word); i++ {
		s := ""
		for j := i; j < len(word); j++ {
			s += string(word[j])
			res += count_vowels(s)
		}
	}
	return res
}

func count_vowels(s string) int64 {
	n := len(s)
	cnt := int64(0)
	for i := 0; i < n; i++ {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
			cnt++
		}
	}
	return cnt
}

//------------------brute force---------------------

//------------------efficient way---------------------
func vowels_of_all_substrings_eff(word string) int64 {
	n := len(word)
	res := int64(0)
	for i := 0; i < n; i++ {
		ch := word[i]
		if is_vowel(ch) {
			res += int64(i+1) * int64(n-i)
		}
	}
	return res
}

func is_vowel(ch byte) bool {
	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
}

//------------------efficient way---------------------

func main() {
	fmt.Println(vowels_of_all_substrings_eff("aba"))
}
