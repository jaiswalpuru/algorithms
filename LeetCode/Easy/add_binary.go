/*
Given two binary strings a and b, return their sum as a binary string.
*/

package main

import "fmt"

func add(str1, str2 string) string {
	res := ""
	c := 0
	var val byte
	i, j := len(str1)-1, len(str2)-1
	for i >= 0 && j >= 0 {
		val, c = bin_add(str1[i], str2[j], c)
		res = string(val) + res
		i--
		j--
	}
	if i >= 0 {
		for i >= 0 {
			val, c = bin_add(str1[i], '0', c)
			res = string(val) + res
			i--
		}
	}
	if j >= 0 {
		for j >= 0 {
			val, c = bin_add('0', str2[j], c)
			res = string(val) + res
			j--
		}
	}
	if i < 0 && j < 0 && c != 0 {
		res = "1" + res
	}
	return res
}

func bin_add(a, b byte, c int) (byte, int) {
	if a == '0' && b == '0' {
		if c == 0 {
			return '0', 0
		} else {
			return '1', 0
		}
	} else if (a == '1' && b == '0') || (a == '0' && b == '1') {
		if c == 0 {
			return '1', 0
		} else {
			return '0', 1
		}
	} else {
		if c == 0 {
			return '0', 1
		} else {
			return '1', 1
		}
	}
}

func main() {
	fmt.Println(add("11", "1"))
}
