package main

import "fmt"

func add_digits(n int) int {
	if n == 0 {
		return 0
	}
	if n%9 == 0 {
		return 9
	}
	return n % 9
}

func main() {
	fmt.Println(add_digits(38))
}
