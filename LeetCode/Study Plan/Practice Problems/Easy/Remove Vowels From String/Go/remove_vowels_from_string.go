package main

import "fmt"

func remove_vowels_from_string(s string) string {
	str := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
			continue
		}
		str = append(str, s[i])
	}
	return string(str)
}

func main() {
	fmt.Println(remove_vowels_from_string("leetcodeisacommunityforcoders"))
}
