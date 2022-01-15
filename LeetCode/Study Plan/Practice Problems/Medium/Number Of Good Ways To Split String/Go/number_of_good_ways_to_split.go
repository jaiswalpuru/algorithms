package main

import "fmt"

//--------------------bruteforce---------------------------------------
func split_bf(s string) int {
	res := 0
	n := len(s)

	for split_ind := 0; split_ind < n; split_ind++ {
		left := make(map[byte]int)
		right := make(map[byte]int)

		for i := 0; i <= split_ind; i++ {
			left[s[i]]++
		}
		for i := split_ind + 1; i < n; i++ {
			right[s[i]]++
		}
		if len(left) == len(right) {
			res++
		}
	}
	return res
}

//---------------------------------------------------------------------

//----------------------------------some optimization to above code---------------------
func split(s string) int {

	n := len(s)
	left := make(map[string]int)
	right := make(map[string]int)
	left_visited := make(map[byte]struct{})
	right_visited := make(map[byte]struct{})
	left[string(s[0])]++
	left_visited[s[0]] = struct{}{}
	right[string(s[n-1])]++
	right_visited[s[n-1]] = struct{}{}
	res := 0

	for i := 1; i < n; i++ {
		if _, ok := left_visited[s[i]]; ok {
			left[s[:i+1]] = left[s[:i]]
		} else {
			left[s[:i+1]] = left[s[:i]] + 1
			left_visited[s[i]] = struct{}{}
		}

		if _, ok := right_visited[s[n-1-i]]; ok {
			right[s[n-1-i:]] = right[s[n-i:]]
		} else {
			right[s[n-1-i:]] = right[s[n-i:]] + 1
			right_visited[s[n-1-i]] = struct{}{}
		}
	}

	for i := 0; i < n; i++ {
		if left[s[:i+1]] == right[s[i+1:]] {
			res++
		}
	}

	return res
}

//----------------------------------------------------------------------------------------

func main() {
	fmt.Println(split("acbadbaada"))
}
