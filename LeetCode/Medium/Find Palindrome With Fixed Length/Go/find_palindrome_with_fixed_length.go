package main

import (
	"fmt"
	"math"
	"strconv"
)

func find_palindrome_with_fixed_length(queries []int, int_length int) []int64 {
	res := []int64{}
	n := len(queries)

	for i := 0; i < n; i++ {
		res = append(res, nth_palindrome(queries[i], int_length))
	}
	return res
}

func nth_palindrome(n, k int) int64 {
	t := 0
	if k%2 == 1 {
		t = k / 2
	} else {
		t = k/2 - 1
	}

	palin := int64(math.Pow(10, float64(t)))
	palin += int64(n - 1)

	str := strconv.Itoa(int(palin))
	if k%2 == 1 {
		palin /= 10
	}
	rev := ""
	for palin > 0 {
		rev = rev + strconv.Itoa(int(palin%10))
		palin = palin / 10
	}

	if rev != "0" {
		str += rev
	}

	val, _ := strconv.Atoi(str)

	if k < len(str) {
		return -1
	}

	return int64(val)
}

func main() {
	fmt.Println(find_palindrome_with_fixed_length([]int{1, 2, 3, 4, 5, 90}, 3))
}
