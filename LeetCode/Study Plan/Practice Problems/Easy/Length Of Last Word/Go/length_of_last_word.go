package main

import "fmt"

func length_of_last_word(s string) int {
	ans := 0
	temp := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if temp != 0 {
				ans = temp
				temp = 0
			}
		} else {
			temp++
		}
	}
	if temp > 0 {
		ans = temp
	}
	return ans
}

func main() {
	fmt.Println(length_of_last_word("Hello World"))
}
