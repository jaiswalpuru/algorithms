package main

import "fmt"

func match(p, q string, i, j int) bool {
	if i == len(p) && j == len(q) {
		return true
	}

	if i == len(p) {
		return false
	}

	if j < len(q) && p[i] == q[j] {
		return match(p, q, i+1, j+1)
	}

	if p[i] >= 'a' && p[i] <= 'z' {
		return match(p, q, i+1, j)
	}
	return false
}

func camel_case_matching(str []string, pattern string) []bool {
	n := len(str)
	res := make([]bool, n)

	for i := 0; i < n; i++ {
		res[i] = match(str[i], pattern, 0, 0)
	}

	return res
}

func main() {
	fmt.Println(camel_case_matching([]string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}, "FB"))
}
