package main

import "fmt"
import "math"

func rearrange_characters_to_make_target_string(s, target string) int {
	mp := make(map[byte]int)
	str := []byte(s)
	for i := 0; i < len(str); i++ {
		mp[str[i]]++
	}

	min_val := math.MaxInt64
	tar := []byte(target)
	vi := make(map[byte]int)
	for i := 0; i < len(tar); i++ {
		vi[tar[i]]++
	}

	for k, v := range vi {
		min_val = min(min_val, mp[k]/v)
	}
	if min_val == math.MaxInt64 {
		return 0
	}
	return min_val
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(rearrange_characters_to_make_target_string("ilovecodingonleetcode", "code"))
}
