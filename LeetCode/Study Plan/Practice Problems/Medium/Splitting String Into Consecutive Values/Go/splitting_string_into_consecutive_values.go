package main

import "fmt"
import "strconv"

func splitting_string_into_consecutive_value(s string) bool {
	n := len(s)
	if n == 1 {
		return false
	}
	res := false
	_splitting_string_into_consecutive_value(s, 0, []int{}, &res)
	return res
}

func is_safe(set []int) bool {
	if len(set) > 1 {
		n1 := set[len(set)-2]
		n2 := set[len(set)-1]
		return n1-n2 == 1
	}
	return true
}

func _splitting_string_into_consecutive_value(s string, start int, set []int, res *bool) {
	if start == len(s) {
		if len(set) >= 2 {
			*res = true
		}
		return
	}

	for i := start; i < len(s); i++ {
		num, _ := strconv.Atoi(s[start : i+1])
		if num == 0 && i != len(s)-1 {
			continue
		}
		temp := append(set, num)
		if is_safe(temp) {
			_splitting_string_into_consecutive_value(s, i+1, temp, res)
		}
		temp = temp[:len(temp)-1]
	}
}

func main() {
	fmt.Println(splitting_string_into_consecutive_value("1234"))
	fmt.Println(splitting_string_into_consecutive_value("050043"))
}
