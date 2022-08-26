package main

import (
	"fmt"
	"reflect"
)

func count_digits(n int) map[int]int {
	mp := make(map[int]int)
	for n > 0 {
		mp[n%10]++
		n /= 10
	}
	return mp
}

func reordered_power_of_two(n int) bool {
	arr := count_digits(n)
	for i := 0; i < 31; i++ {
		curr := count_digits(1 << i)
		if reflect.DeepEqual(arr, curr) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(reordered_power_of_two(10))
	fmt.Println(reordered_power_of_two(16))
}
