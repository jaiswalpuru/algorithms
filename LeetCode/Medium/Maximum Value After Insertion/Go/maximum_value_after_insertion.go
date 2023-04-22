package main

import (
	"fmt"
	"strconv"
)

func maximum_value_after_insertion(n string, x int) string {
	l := len(n) - 1
	pos := l + 1
	if n[0] == '-' {
		for i := l; i >= 1; i-- {
			if int(n[i]-'0') > x {
				pos = i
			}
		}
	} else {
		for i := l; i >= 0; i-- {
			if int(n[i]-'0') < x {
				pos = i
			}
		}
	}
	return n[:pos] + strconv.Itoa(x) + n[pos:]
}

func main() {
	fmt.Println(maximum_value_after_insertion("99", 9))
}
