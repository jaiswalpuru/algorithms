/*
Given a string s consisting of some words separated by some number of spaces, return the length of the
last word in the string.

A word is a maximal substring consisting of non-space characters only.
*/

package main

import (
	"fmt"
	"strings"
)

func length(str string) int {
	s := strings.Split(strings.TrimSpace(str), " ")
	return len(s[len(s)-1])
}

func main() {
	fmt.Println(length("Hello World"))
	fmt.Println(length("   fly me   to   the moon  "))
	fmt.Println(length("luffy is still joyboy"))
}
