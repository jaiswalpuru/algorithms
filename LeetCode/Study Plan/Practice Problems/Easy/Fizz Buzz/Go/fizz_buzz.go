package main

import (
	"fmt"
	"strconv"
)

func fizz_buzz(n int) []string {
	res := make([]string, n+1)
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			res[i] = "FizzBuzz"
		} else if i%5 == 0 {
			res[i] = "Buzz"
		} else if i%3 == 0 {
			res[i] = "Fizz"
		} else {
			res[i] = strconv.Itoa(i)
		}
	}
	return res[1:]
}

func main() {
	fmt.Println(fizz_buzz(15))
}
