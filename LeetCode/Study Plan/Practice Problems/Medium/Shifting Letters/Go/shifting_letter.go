package main

import "fmt"

func shifting_letter(s string, shifts []int) string {
	s_rune := []rune(s)
	n := len(s_rune)
	if n == 1 {
		shifts[0] = shifts[0] % 26
	}
	for i := n - 2; i >= 0; i-- {
		shifts[i] = (shifts[i+1] + shifts[i]) % 26
	}
	for i := 0; i < n; i++ {
		s_rune[i] = rune(int(s_rune[i]) % 97)
		s_rune[i] = rune((int(s_rune[i])+shifts[i])%26) + 'a'
	}
	return string(s_rune)
}

func main() {
	fmt.Println(shifting_letter("zzz", []int{1, 2, 3}))
}
