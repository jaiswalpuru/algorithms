package main

import "fmt"

func divide_the_string_into_groups_of_size_k(s string, k int, fill byte) []string {
	if len(s)%k != 0 {
		val := k - len(s)%k
		for i := 0; i < val; i++ {
			s += string(fill)
		}
	}
	res := []string{}
	for i := 0; i < len(s)-k+1; i += k {
		res = append(res, s[i:i+k])
	}
	return res
}

func main() {
	fmt.Println(divide_the_string_into_groups_of_size_k("abcdefghi", 3, 'x'))
}
