package main

import "fmt"

func swap_adjacent_in_lr_string(start, end string) bool {
	j := 1
	s, e := []byte(start), []byte(end)
	for i := 0; i < len(start); i++ {
		if s[i] == e[i] {
			continue
		}
		if (e[i] == 'X' && s[i] == 'R') || (e[i] == 'L' && s[i] == 'X') {
			j = max(j, i+1)
			for j < len(s) && s[j] == s[i] {
				j++
			}
			if j == len(s) || s[j] != e[i] {
				return false
			}
			s[j] = s[i]
		} else {
			return false
		}
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(swap_adjacent_in_lr_string("RXXLRXRXL", "XRLXXRRLX"))
}
