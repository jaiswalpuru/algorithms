package main

import "fmt"
import "strconv"

func largest_3_same_digit_number_in_string(str string) string {
	n := len(str)
	s := string(str[0])
	res, f := 0, false
	for i := 1; i < n; i++ {
		if str[i] == str[i-1] {
			s += string(str[i])
		} else {
			if len(s) >= 3 {
				val, _ := strconv.Atoi(s[0:3])
				if val >= res {
					res = val
					f = true
				}
			}
			s = string(str[i])
		}
	}
	if len(s) >= 3 {
		val, _ := strconv.Atoi(s[0:3])
		if val >= res {
			res = val
			f = true
		}
	}
	if res == 0 {
		if f {
			return "000"
		}
		return ""
	}
	temp := strconv.Itoa(res)
	return temp
}

func main() {
	fmt.Println(largest_3_same_digit_number_in_string("6777133339"))
}
