package main

import (
	"fmt"
	"math"
)

func maximum_number_of_balloons(text string) int {
	char_cnt := make([]int, 26)
	for i := 0; i < len(text); i++ {
		char_cnt[int(text[i])-'a']++
	}
	str := "balloon"
	res := math.MaxInt64
	for i := 0; i < len(str); i++ {
		if str[i] == 'l' {
			res = min(res, char_cnt[str[i]-'a']/2)
		} else if str[i] == 'o' {
			res = min(res, char_cnt[str[i]-'a']/2)
		} else {
			res = min(res, char_cnt[str[i]-'a'])
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maximum_number_of_balloons("nlaebolko"))
}
