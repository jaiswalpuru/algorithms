package main

import "fmt"

func parens(n int) []string {
	res := []string{}
	s := make([]byte, n*2)
	recur(n, n, s, &res, 0)
	return res
}

func recur(n, m int, set []byte, res *[]string, ind int) {
	if n < 0 || m < n {
		return
	}
	if n == 0 && m == 0 {
		*res = append(*res, string(set))
		return
	}
	set[ind] = '('
	recur(n-1, m, set, res, ind+1)
	set[ind] = ')'
	recur(n, m-1, set, res, ind+1)
}

func main() {
	fmt.Println(parens(3))
}
