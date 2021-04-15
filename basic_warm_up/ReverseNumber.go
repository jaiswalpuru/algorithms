package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func reverseNumber(n int) int {

	temp := 0

	for n > 0 {
		t := n % 10
		temp = (temp + t) * 10
		n /= 10
	}

	return temp

}

func main() {

	var (
		t, n int
		sc   = bufio.NewScanner(os.Stdin)
	)

	sc.Scan()
	t, _ = strconv.Atoi(sc.Text())

	for t > 0 {
		sc.Scan()
		n, _ = strconv.Atoi(sc.Text())
		res := reverseNumber(n)
		fmt.Println(res / 10)
		t--
	}

}
