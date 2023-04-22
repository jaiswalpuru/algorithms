package main

import "fmt"
import "strconv"

func get_sum(str string) int {
	n := len(str)
	sum := 0
	for i := 0; i < n; i++ {
		sum += int(str[i] - '0')
	}
	return sum
}

func digit_sum(s string, k int) string {
	res := ""
	for len(s) > k {
		n := len(s)
		t := ""
		for i := 0; i < n; i++ {
			if len(t) < k {
				t += string(s[i])
			} else {
				res += strconv.Itoa(get_sum(t))
				t = string(s[i])
			}
		}
		if t != "" {
			res += strconv.Itoa(get_sum(t))
		}
		s = res
		res = ""
		t = ""
	}
	return s
}

func main() {
	fmt.Println(digit_sum("11111222223", 3))
}
