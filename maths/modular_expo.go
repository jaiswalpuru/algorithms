package main

import "fmt"

//Prime -> 10^9 + 7 use modulo

var (
	m = 1000000007
)

func modularExpo(a, n int) int {
	if n == 0 {
		return 1
	}

	if n%2 == 0 {
		halfAns := modularExpo(a, n/2)
		return (halfAns * halfAns) % m
	} else {
		halfAns := modularExpo(a, n/2)
		halfAns = (halfAns * halfAns) % m //(a^n/2 * a^n/2)%m
		halfAns = (halfAns * a) % m
		return halfAns
	}
}

func main() {
	fmt.Println(modularExpo(10, 10))
}
