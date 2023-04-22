package main

import (
	"fmt"
	"math"
	"strconv"
)

func split_array_into_fibonacci_sequnce(num string) []int {
	n := len(num)
	if n < 3 {
		return nil
	}

	res := []int{}
	if backtrack(num, 0, &res) {
		return res
	}
	return nil
}

func backtrack(num string, idx int, res *[]int) bool {
	if idx == len(num) {
		return len(*res) > 2
	}

	s := ""
	for i := idx; i < len(num); i++ {
		if num[idx] == '0' && i != idx {
			return false
		}
		s += string(num[i])
		val, _ := strconv.Atoi(s)
		if val > math.MaxInt32 {
			break
		}
		if len(*res) < 2 || (val-(*res)[len(*res)-2] == (*res)[len(*res)-1]) {
			*res = append(*res, val)
			if backtrack(num, i+1, res) {
				return true
			}
			*res = (*res)[:len(*res)-1]
		}
	}
	return false
}

func main() {
	fmt.Println(split_array_into_fibonacci_sequnce("539834657215398346785398346991079669377161950407626991734534318677529701785098211336528511"))
}
