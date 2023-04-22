package main

import "fmt"

var val = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

var hash_map = map[int]string{
	1: "I", 4: "IV", 5: "V", 9: "IX", 10: "X", 40: "XL", 50: "L", 90: "XC", 100: "C", 400: "CD", 500: "D", 900: "CM", 1000: "M",
}

func integer_to_roman(num int) string {
	res := ""
	for i := 0; i < len(val) && num > 0; i++ {
		for val[i] <= num {
			num -= val[i]
			res += hash_map[val[i]]
		}
	}
	return res
}

func main() {
	fmt.Println(integer_to_roman(3))
}
