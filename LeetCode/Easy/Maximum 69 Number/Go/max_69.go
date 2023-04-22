package main

import (
	"fmt"
	"strconv"
)

func max_69(n int) int {
	s := strconv.Itoa(n)
	str := []byte(s)
	for i := 0; i < len(str); i++ {
		if str[i] == '6' {
			str[i] = '9'
			break
		}
	}
	val, _ := strconv.Atoi(string(str))
	return val
}

func main() {
	fmt.Println(max_69(9669))
}
