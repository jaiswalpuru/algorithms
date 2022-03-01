package main

import "fmt"

func counting_bits(n int) []int {
	arr := make([]int, n+1)
	arr[0] = 0

	for i := 1; i <= n; i++ {
		val := i
		cnt := 0
		for val > 0 {
			if val&1 == 1 {
				cnt++
			}
			val >>= 1
		}
		arr[i] = cnt
	}
	return arr
}

func main() {
	fmt.Println(counting_bits(2))
}
