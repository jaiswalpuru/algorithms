package main

import "fmt"

func check_palindrome(s string) bool {
	n := len(s)
	if n == 0 || n == 1 {
		return true
	} else if s[0] == s[n-1] {
		return check_palindrome(s[1 : n-1])
	} else {
		return false
	}
}

//---------------------Brute force bad solution------------------------------------------
func max_prod_length(s string) int {
	res := [][]int{}
	_subsequence(s, &res, []int{}, 0)
	ans := 1
	n := len(res)
	for i := 0; i < n; i++ {

		var m1, m2 int

		var p, q []int
		p = res[i]
		m1 = len(p)
		//create the strings
		str1, str2 := "", ""
		for t := 0; t < m1; t++ {
			str1 += string(s[p[t]])
		}
		if !check_palindrome(str1) {
			continue
		}

		mp := make(map[int]struct{})
		for j := 0; j < m1; j++ {
			mp[p[j]] = struct{}{}
		}
		for j := i + 1; j < n; j++ {
			q = res[j]
			// fmt.Print(p, " ")
			// fmt.Println(q)
			m2 = len(q)

			str2 = ""
			for t := 0; t < m2; t++ {
				str2 += string(s[q[t]])
			}
			if !check_palindrome(str2) {
				continue
			}

			f := true
			for k := 0; k < m2; k++ {
				if _, ok := mp[q[k]]; ok {
					f = false
					break
				}
			}
			if f {
				ans = max(ans, m1*m2)
			}
		}

	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func _subsequence(s string, res *[][]int, set []int, start int) {
	if start >= len(s) {
		*res = append(*res, append([]int{}, set...))
		return
	}

	temp := append(set, start)
	_subsequence(s, res, temp, start+1)
	temp = temp[:len(temp)-1]
	_subsequence(s, res, temp, start+1)
}

//------------------------------------------------------------------------------------

//----------------------------------------optimized soln------------------------------
func max_prod_length_optimized(s string) int {
	ans := 1
	_max_prod_length_optimized(s, "", "", 0, &ans)
	return ans
}

func _max_prod_length_optimized(s, s1, s2 string, start int, ans *int) {
	if start == len(s) {
		if check_palindrome(s1) && check_palindrome(s2) {
			*ans = max(*ans, len(s1)*len(s2))
		}
		return
	}
	temp := string(s[start])
	_max_prod_length_optimized(s, s1+temp, s2, start+1, ans)
	_max_prod_length_optimized(s, s1, s2+temp, start+1, ans)
	_max_prod_length_optimized(s, s1, s2, start+1, ans)
}

//------------------------------------------------------------------------------------

func main() {
	fmt.Println(max_prod_length_optimized("yyyyyynyyyy"))
}
