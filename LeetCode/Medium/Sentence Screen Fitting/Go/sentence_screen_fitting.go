package main

import "fmt"

func sentence_screen_fitting(sentence []string, row, col int) int {
	i := 0
	cnt := 0
	c := col
	for i < row {
		k := 0
		for k < len(sentence) && i < row {
			if c >= len(sentence[k]) {
				c -= (len(sentence[k]) + 1)
				k++
			} else {
				i++
				c = col
			}
		}
		if k == len(sentence) {
			cnt++
		}
	}
	return cnt
}

func main() {
	fmt.Println(sentence_screen_fitting([]string{"f", "p", "a"}, 8, 7))
}
