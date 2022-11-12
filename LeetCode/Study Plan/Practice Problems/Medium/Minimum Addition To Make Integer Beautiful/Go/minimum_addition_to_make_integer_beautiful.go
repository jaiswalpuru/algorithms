package main

import "fmt"

func minimum_addition_to_make_integer_beautiful(n int64, target int) int64 {
	mul := int64(1)
	cnt := int64(0)
	for sum_of_digits(n) > int64(target) {
		cnt += mul * (int64(10) - n%10)
		mul *= 10
		n = n/10 + 1
	}
	return cnt
}

func sum_of_digits(n int64) int64 {
	sum := int64(0)
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func main() {
	fmt.Println(minimum_addition_to_make_integer_beautiful(16, 6))
}
