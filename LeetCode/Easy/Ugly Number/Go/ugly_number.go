package main

import "fmt"

func ugly_number(n int) bool {
	for n%2 == 0 {
		n = n / 2
	}
	for n%3 == 0 {
		n = n / 3
	}
	for n%5 == 0 {
		n = n / 5
	}
	return n == 1
}

func main() {
	fmt.Println(ugly_number(6))
}
