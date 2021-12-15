package main

import "fmt"

func finbonacci(n int) int {
	if n <= 1 {
		return n
	}
	return finbonacci(n-1) + finbonacci(n-2)
}

func main() {
	fmt.Println(finbonacci(10))
}
