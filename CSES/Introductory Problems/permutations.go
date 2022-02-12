package main

import "fmt"

func solve(n int) []int {
	if n <= 3 {
		return nil
	}
	res := []int{}
	i, j := 1, n
	turn := 0
	for i < j {
		if turn%2 == 0 {
			res = append(res, i)
			i++
		} else {
			res = append(res, j)
			j--
		}
		turn++
	}
	return res
}

func main() {
	fmt.Println(solve(100))
}
