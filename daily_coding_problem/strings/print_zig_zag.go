/*
Print zig zag form
string str := "thisisazigzag" k=4 (no of rows)
1 t     a     g
2  h   s z   a
3   i i   i z
4    s     g
*/
package main

import (
	"fmt"
	"strings"
)

func is_descending(i, k int) bool {
	maxSpaces := 2*(k-1) - 1
	//adding 1 to maxspace lets say i=6 and maxspace=5, but we need to oscillated between them 6%6=0
	return i%(maxSpaces+1) < k-1
}

func get_space(row_num int, k int, desc bool) int {
	spaces := 0
	maxSpace := (k-1)*2 - 1
	if desc {
		spaces = maxSpace - row_num*2
	} else {
		spaces = maxSpace - (k-1-row_num)*2
	}
	return spaces
}

func print_zig_zag(str string, k int) {
	n := len(str)

	for row := 0; row < k; row++ {

		i := row
		line := []string{}
		for j := 0; j < n; j++ {
			line = append(line, " ")
		}
		for i < n {
			line[i] = string(str[i])
			desc := is_descending(i, k)
			spaces := get_space(row, k, desc)
			i += spaces + 1
		}
		fmt.Println(strings.Join(line, ""))
	}
}

func main() {
	str := "thisisazigzag"
	k := 4
	print_zig_zag(str, k)
}
