package main

import "fmt"

func daily_temperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	st := []int{}
	for i := 0; i < n; i++ {
		for len(st) > 0 && temperatures[i] > temperatures[st[len(st)-1]] {
			res[st[len(st)-1]] = i - st[len(st)-1]
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}
	return res
}

func main() {
	fmt.Println(daily_temperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}
