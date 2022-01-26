package main

import "fmt"

func solve(n int) {
	fmt.Print(n, "->")
	for n > 0 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}
		if n == 1 {
			fmt.Println(n)
			return
		} else {
			fmt.Print(n, "->")
		}
	}
}

func main() {
	solve(3)
}
