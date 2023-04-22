package main

import "fmt"

func decode_the_message(key, message string) string {
	c := 'a'
	n := len(key)
	mp := make(map[byte]rune)
	for i := 0; i < n; i++ {
		if key[i] != ' ' {
			if _, ok := mp[key[i]]; ok {
				continue
			}
			mp[key[i]] = c
			c += 1
		}
	}
	res := ""
	n = len(message)
	for i := 0; i < n; i++ {
		if message[i] == ' ' {
			res += " "
			continue
		}
		res += string(mp[message[i]])
	}
	return res
}

func main() {
	fmt.Println(decode_the_message("the quick brown fox jumps over the lazy dog", "vkbs bs t suepuv"))
}
