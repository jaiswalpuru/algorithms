package main

import "fmt"

func hammingWeight(num uint32) int {
	n := uint32(1)
	cnt := 0
	for num > 0 {
		if (num & n) == 1 {
			cnt++
		}
		num = num >> 1
	}
	return cnt
}

func main() {
	fmt.Println(hammingWeight(00000000000000000000000000001011))
}
