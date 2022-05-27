package main

import "fmt"

func number_of_steps_to_reduce_number_to_zero(num int) int {
	steps := 0
	for num > 0 {
		if num%2 == 0 {
			num /= 2
		}else {
			num -= 1
		}
		steps++
	}
	return steps
}

func main() {
	fmt.Println(number_of_steps_to_reduce_number_to_zero(14))
}