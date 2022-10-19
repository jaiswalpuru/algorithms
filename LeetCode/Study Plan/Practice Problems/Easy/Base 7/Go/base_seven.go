package main

import (
	"fmt"
	"strconv"
)

func base_seven(n int) string {
	str := ""
	neg := false
	if n == 0 {
		return "0"
	}
	if n < 0 {
		neg = true
		n = n * -1
	}
	for n > 0 {
		str = strconv.Itoa(n%7) + str
		n = n / 7
	}
	if neg {
		return "-" + str
	}
	return str
}

func main() {
	fmt.Println(base_seven(100))
}
