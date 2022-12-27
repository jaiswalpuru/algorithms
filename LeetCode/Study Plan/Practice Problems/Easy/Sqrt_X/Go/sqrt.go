package main

import "fmt"

func sqrt(x int) int {
	l, r := 0, x
	val := 0
	for l <= r {
		mid := (l + r) / 2
		if mid*mid <= x {
			val = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return val
}

func main() {
	fmt.Println(sqrt(4))
}
