package main

import "fmt"

func sum_of_number_and_its_reverse(num int) bool {
	i := 0
	for i <= num {
		val := i
		rev_val := reverse(val)
		if val+rev_val == num {
			return true
		}
		i++
	}
	return false
}

func reverse(a int) int {
	rev := 0
	for a > 0 {
		temp := a % 10
		rev = rev*10 + temp
		a = a / 10
	}
	return rev
}

func main() {
	fmt.Println(sum_of_number_and_its_reverse(443))
}
