package main

import "fmt"

func parens(n int) []string {
	res := []string{}
	recur(n, n, []byte{}, &res)
	return res
}

func recur(n, m int, set []byte, res *[]string) {
	if n < 0 || m < n {
		return
	}
	if n == 0 && m == 0 {
		*res = append(*res, string(set))
		return
	}
	set = append(set, '(')
	recur(n-1, m, set, res)
	set = set[:len(set)-1]
	set = append(set, ')')
	recur(n, m-1, set, res)
}

func main() {
	fmt.Println(parens(3))
}
