package main

import "fmt"

func percent_of_letter_in_string(s string, b byte) int {
	cnt := 0
	n := len(s)
	str := []byte(s)
	for i := 0; i < n; i++ {
		if b == str[i] {
			cnt++
		}
	}
	return 100 * cnt / n
}

func main() {
	fmt.Println(percent_of_letter_in_string("foobar", 'o'))
}
