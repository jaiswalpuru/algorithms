package main

import "fmt"

func check_if_sentence_is_panagram(s string) bool {
	mp := make(map[byte]struct{})
	n := len(s)
	for i := 0; i < n; i++ {
		mp[s[i]] = struct{}{}
	}
	return len(mp) == 26
}

func main() {
	fmt.Println(check_if_sentence_is_panagram("thequickbrownfoxjumpsoverthelazydog"))
}
