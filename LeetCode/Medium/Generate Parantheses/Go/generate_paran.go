package main

import "fmt"

func generate_parantheses(n int) []string {
	res := []string{}
	_generate(n, "", &res, 0, 0)
	return res
}

func _generate(n int, s string, res *[]string, l, r int) {
	if len(s) == 2*n {
		*res = append(*res, s)
		return
	}

	if l < n {
		s += "("
		_generate(n, s, res, l+1, r)
		s = s[:len(s)-1]
	}
	if l > r {
		s += ")"
		_generate(n, s, res, l, r+1)
		s = s[:len(s)-1]
	}
}

func main() {
	fmt.Println(generate_parantheses(3))
}
