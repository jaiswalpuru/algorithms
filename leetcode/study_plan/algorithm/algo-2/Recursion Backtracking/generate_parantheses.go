/*
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
*/

/*--------------------------------------------------
Solution:
Brute Force :


-----------------------------------------------------*/

package main

import "fmt"

var ans = []string{}

func is_paran_valid(str string) bool {
	bal := 0
	for i := 0; i < len(str); i++ {
		if str[i] == '(' {
			bal++
		} else {
			bal--
		}
		if bal < 0 {
			return false
		}
	}
	return bal == 0
}

func generate_paran_brute_force(arr []byte, n int) {
	if len(arr) == 2*n {
		if is_paran_valid(string(arr)) {
			ans = append(ans, string(arr))
		}
	} else {
		arr = append(arr, '(')
		generate_paran_brute_force(arr, n)
		arr = arr[:len(arr)-1]
		arr = append(arr, ')')
		generate_paran_brute_force(arr, n)
		arr = arr[:len(arr)-1]
	}
}

func generate_parantheses(n int) {
	generate_paran_brute_force([]byte{}, n)
}

func generate_paran_optimized(n int) {
	backtrack([]byte{}, 0, 0, n)
}

func backtrack(arr []byte, left, right, n int) {
	if len(arr) == 2*n {
		ans = append(ans, string(arr))
		return
	}

	if left < n {
		arr = append(arr, '(')
		backtrack(arr, left+1, right, n)
		arr = arr[:len(arr)-1]
	}
	if left > right {
		arr = append(arr, ')')
		backtrack(arr, left, right+1, n)
		arr = arr[:len(arr)-1]
	}
}

func main() {
	generate_paran_optimized(3)
	fmt.Println(ans)
}
