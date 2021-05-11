package main

import (
	"fmt"
	"strconv"
)

func encode(str string) string {
	mp := make(map[rune]int)
	sRune := []rune(str)
	enc := ""
	mp[sRune[0]]++

	for i := 1; i < len(sRune); i++ {
		if sRune[i-1] != sRune[i] {
			enc += strconv.Itoa(mp[sRune[i-1]]) + string(sRune[i-1])
			mp[sRune[i-1]] = 0
		}
		mp[sRune[i]]++
	}
	enc += strconv.Itoa(mp[sRune[len(sRune)-1]]) + string(sRune[len(sRune)-1])
	return enc
}

func decode(str string) string {
	dec := ""
	sRune := []rune(str)
	temp := ""
	for i := 0; i < len(sRune); i++ {
		if sRune[i] >= '0' && sRune[i] <= '9' {
			temp += string(sRune[i])
		} else {
			n, _ := strconv.Atoi(temp)
			for j := 0; j < n; j++ {
				dec += string(sRune[i])
			}
			temp = ""
		}
	}
	return dec
}

func main() {
	str := "AAAABBBCCDAA"
	fmt.Println(encode(str))
	fmt.Println(decode(encode(str)))
}
