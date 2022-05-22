package main

import "fmt"

//-----------------------brute force--------------------------
func is_palindrome(s string) bool {
	n := len(s)
	if n == 1 {
		return true
	}
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}

func palindromic_substrings_brute(s string) int {
	cnt := 0
	n := len(s)
	for i:=0; i<n; i++{
    	for j:=i;j<n;j++{
    		if is_palindrome(s[i:j+1]) {
    			cnt++
    		}
    	}
	}
	return cnt
}
//-----------------------brute force--------------------------


//-----------------------memoization--------------------------
func palindromic_substrings(s string) int {
	cnt := 0
	n := len(s)
	memo := make([][]int, n)
	for i:=0;i<n;i++{
		memo[i] = make([]int, n)
		for j:=0;j<n;j++{
			memo[i][j] = -1
		}
	}

	for i:=0; i<n; i++{
    	for j:=i;j<n;j++{
    		cnt += is_palindrome_memo(i,j, s, &memo)
    	}
	}
	return cnt
}

func is_palindrome_memo(i,j int, s string, memo *[][]int) int {
	if (*memo)[i][j] != -1 {
		return (*memo)[i][j]
	}

	if i>=j {
		(*memo)[i][j] = 1
		return (*memo)[i][j]
	}

	if s[i]!=s[j] {
		(*memo)[i][j] = 0
		return(*memo)[i][j]
	}
	
	if is_palindrome_memo(i+1, j-1, s, memo) == 1 {
		(*memo)[i][j] = 1
		return (*memo)[i][j]
	}else {
		(*memo)[i][j] = 0
		return (*memo)[i][j]
	}
}
//-----------------------memoization--------------------------

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(palindromic_substrings("aaa"))
}
