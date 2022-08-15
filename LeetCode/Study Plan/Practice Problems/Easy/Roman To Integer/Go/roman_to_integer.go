package main

import "fmt"

func roman_to_integer(s string) int {
	mp := map[byte]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000,
	}
	n := len(s)
	res := mp[s[n-1]]
	for i := n - 2; i >= 0; i-- {
		if mp[s[i]] >= mp[s[i+1]] {
			res = res + mp[s[i]]
		} else {
			res = res - mp[s[i]]
		}
	}
	return res
}

func main() {
	fmt.Println(roman_to_integer("MCMXCIV"))
}
