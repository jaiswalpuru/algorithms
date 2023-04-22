package main

import "fmt"

func find_pivot_integer(n int) int {
	for i := 1; i <= n; i++ {
		first := i * (i + 1) / 2
		total := n*(n+1)/2 - first + i
		if first == total {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(find_pivot_integer(8))
}
