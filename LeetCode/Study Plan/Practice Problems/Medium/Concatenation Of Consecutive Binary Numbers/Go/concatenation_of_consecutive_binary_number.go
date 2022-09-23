package main

import (
	"fmt"
	"strconv"
)

var mod = int(1e9 + 7)

//-------------------brute force-------------------------
func concatenation_of_consecutive_binary_number(n int) int {
	_n := int64(n)
	str := ""
	for i := int64(1); i <= _n; i++ {
		str += strconv.FormatInt(i, 2)
	}
	val := 0
	j := 0
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] != '0' {
			val += fast_expo(2, j) % mod
		}
		j++
	}
	return val % mod
}

func fast_expo(a, n int) int {
	if n == 0 {
		return 1
	}
	if n%2 == 0 {
		half := fast_expo(a, n/2)
		return half % mod * half % mod
	} else {
		half := fast_expo(a, n/2)
		return half % mod * half % mod * a
	}
}

//-------------------brute force-------------------------

//-------------------better approach-------------------------
func concatenation_of_consecutive_binary_number_one(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		str := strconv.FormatInt(int64(i), 2)
		for j := 0; j < len(str); j++ {
			res = (res*2 + int(str[j]-'0')) % mod
		}
	}
	return res
}

//-------------------better approach-------------------------

//-------------------mathematical approach-------------------
func concatenation_of_consecutive_binary_number_two(n int) int {
	res := 0
	length := 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			length++
		}
		res = ((res << length) | i) % mod
	}
	return res
}

//-------------------mathematical approach-------------------

func main() {
	fmt.Println(concatenation_of_consecutive_binary_number_two(3))
}
