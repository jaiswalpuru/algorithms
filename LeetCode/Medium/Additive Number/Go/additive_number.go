package main

import "fmt"
import "strconv"

func additive_number(num string) bool {
	n := len(num)
	if n < 3 {
		return false
	}
	res := false
	_additive_number(num, 0, []string{}, &res)
	return res
}

func is_additive(set []string) bool {
	if len(set) >= 3 {
		n1, _ := strconv.Atoi(set[len(set)-3])
		n2, _ := strconv.Atoi(set[len(set)-2])
		n3, _ := strconv.Atoi(set[len(set)-1])
		return n1+n2 == n3
	}
	return true
}

func _additive_number(num string, start int, set []string, res *bool) {
	if start == len(num) {
		if len(set) >= 3 {
			*res = true
		}
		return
	}

	for i := start; i < len(num); i++ {
		n := num[start : i+1]
		if len(n) > 1 && n[0] == '0' {
			continue
		}
		temp := append(set, n)
		if is_additive(temp) {
			_additive_number(num, i+1, temp, res)
		}
		temp = temp[:len(temp)-1]
	}
}

func main() {
	fmt.Println(additive_number("199100199"))
}
