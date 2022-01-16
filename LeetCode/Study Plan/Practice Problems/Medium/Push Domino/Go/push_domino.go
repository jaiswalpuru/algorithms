package main

import "fmt"

func push_domino(s string) string {
	n := len(s)
	index := make([]int, n+2)
	symbols := make([]byte, n+2)

	l := 1
	index[0] = -1
	symbols[0] = 'L'

	for i := 0; i < n; i++ {
		if s[i] != '.' {
			index[l] = i
			symbols[l] = s[i]
			l++
		}
	}
	index[l] = n
	symbols[l] = 'R'
	l++
	char_arr := []byte(s)
	for ind := 0; ind < l-1; ind++ {
		i, j := index[ind], index[ind+1]
		p, q := symbols[ind], symbols[ind+1]

		if p == q {
			for k := i + 1; k < j; k++ {
				char_arr[k] = p
			}
		} else if p > q {
			for k := i + 1; k < j; k++ {
				if k-i == j-k {
					char_arr[k] = '.'
				} else if k-i < j-k {
					char_arr[k] = 'R'
				} else {
					char_arr[k] = 'L'
				}
			}
		}
	}
	return string(char_arr)
}

func main() {
	fmt.Println(push_domino(".L.R...LR..L.."))
}
