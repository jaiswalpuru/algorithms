package main

import (
	"fmt"
)

//-----------------------------Brute force TLE-----------------------------
func can_convert(s1, s2 string, k int) bool {

	n, m := len(s1), len(s2)
	if n != m {
		return false
	}

	b1, b2 := []byte(s1), []byte(s2)
	for i := 1; i <= k; i++ {
		for j := 0; j < n; j++ {
			if b1[j] != b2[j] {
				if (int(b1[j])-97+i)%26+97 == int(b2[j]) {
					b1[j] = b2[j]
					break
				}
			}
		}
	}
	if string(b1) == string(b2) {
		return true
	}
	return false
}

//-----------------------------Brute force TLE-----------------------------

//-----------------------------Optimized solution-----------------------------
func can_convert_optimized(s1, s2 string, k int) bool {
	n, m := len(s1), len(s2)
	if n != m {
		return false
	}

	var shifs [26]int

	for i := 0; i < n; i++ {
		p, q := int(s1[i]-'a'), int(s2[i]-'a')
		if p != q {
			move := (q - p + 26) % 26
			shifs[move]++
			if move+(shifs[move]-1)*26 > k {
				return false
			}
		}
	}
	return true
}

//-----------------------------Optimized solution-----------------------------
func main() {
	// fmt.Println(can_convert_optimized("iqssxdlb", "dyuqrwyr", 40))
	fmt.Println(can_convert_optimized("aab", "bbb", 27))
}
