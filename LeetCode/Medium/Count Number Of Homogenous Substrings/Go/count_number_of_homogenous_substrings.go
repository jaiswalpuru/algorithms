package main

import "fmt"

var mod = int(1e9 + 7)

func count_number_of_homogenous_substrings(s string) int {
	res := 0
	cnt := 0
	for i := 0; i < len(s); i++ {
		if i == 0 {
			cnt++
			continue
		}
		if s[i] == s[i-1] {
			cnt++
		} else {
			res += ((cnt * (cnt + 1)) % mod) / 2
			cnt = 1
		}
	}
	res += ((cnt * (cnt + 1)) % mod) / 2
	return res
}

func main() {
	fmt.Println(count_number_of_homogenous_substrings("abbcccaa"))
}
