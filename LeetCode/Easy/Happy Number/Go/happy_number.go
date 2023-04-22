package main

import "fmt"

func happy_number(n int) bool {
	visited := make(map[int]bool)
	for {
		visited[n] = true
		num := 0
		for n > 0 {
			num += (n % 10) * (n % 10)
			n = n / 10
		}
		n = num
		if n == 1 {
			return true
		}
		if visited[n] {
			return false
		}
	}
	return false
}

func main() {
	fmt.Println(happy_number(19))
}
