package main

import "fmt"

func smallest_string(n, k int) string {
	alpha_map := make(map[int]byte)
	start := byte('a')
	for i:=1; i<=26; i++ {
		alpha_map[i] = start
		start++
	}

	str := make([]byte, n+1)

	char := 26
	temp_pos := n
	for i:=n; i>0; i-- {
		if k - char >= temp_pos-1 {
			str[i] = alpha_map[char]
		} else {
			for k-char < temp_pos-1 {
				char--
			}
			str[i] = alpha_map[char]
		}
		k-=char
		temp_pos--
	}

	return string(str[1:])
}

func main() {
	fmt.Println(smallest_string(3, 27))
}