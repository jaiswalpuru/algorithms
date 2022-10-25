package main

import (
	"fmt"
	"strings"
)

func check_if_two_strings_are_equivalent(s1, s2 []string) bool {
	return strings.Join(s1, "") == strings.Join(s2, "")
}

func main() {
	fmt.Println(check_if_two_strings_are_equivalent([]string{"ab", "c"}, []string{"a", "bc"}))
}
