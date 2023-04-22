package main

import "fmt"

//-----------------------worst brute force ever-------------------
func count_vowels_permutation(n int) int {
	res := []string{"a", "e", "i", "o", "u"}
	if n == 1 {
		return len(res)
	}
	temp := []string{}
	for i := 1; i < n; i++ {
		for i := 0; i < len(res); i++ {
			update_res(res[i], &temp)
		}
		res = temp
		temp = []string{}
	}
	return len(res)
}

func update_res(r string, temp *[]string) {
	char := r[len(r)-1]
	if char == 'a' {
		*temp = append(*temp, r+"e")
	} else if char == 'e' {
		*temp = append(*temp, r+"a")
		*temp = append(*temp, r+"i")
	} else if char == 'i' {
		*temp = append(*temp, r+"a")
		*temp = append(*temp, r+"e")
		*temp = append(*temp, r+"o")
		*temp = append(*temp, r+"u")
	} else if char == 'o' {
		*temp = append(*temp, r+"i")
		*temp = append(*temp, r+"u")
	} else {
		*temp = append(*temp, r+"a")
	}
}

//-----------------------worst brute force ever-------------------

//----------------------a better brute force but still bad TLE----------
//a:0 e:1 i:2 o:3 u:4
var mod = int(1e9 + 7)

func count_vowels_permutation_recursion(n int) int {
	res := 0
	for i := 0; i < 5; i++ {
		res = (res + _recur(n-1, i)) % mod
	}
	return res
}

func _recur(i int, vowel int) int {
	if i == 0 {
		return 1
	}
	if vowel == 0 {
		return (_recur(i-1, 1) + _recur(i-1, 2) + _recur(i-1, 4)) % mod
	} else if vowel == 1 {
		return (_recur(i-1, 0) + _recur(i-1, 2)) % mod
	} else if vowel == 2 {
		return (_recur(i-1, 1) + _recur(i-1, 3)) % mod
	} else if vowel == 3 {
		return (_recur(i-1, 2)) % mod
	} else {
		return (_recur(i-1, 2) + _recur(i-1, 3)) % mod
	}
}

//----------------------a better brute force but still bad TLE----------

//----------------------memoization------------------
func count_vowel_memo(n int) int {
	if n == 1 {
		return 5
	}
	memo := make([][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, 5)
	}
	res := 0
	for i := 0; i < 5; i++ {
		res = (res + _memo(n-1, i, &memo)) % mod
	}
	return res
}

func _memo(i int, vowel int, memo *[][]int) int {
	if (*memo)[i][vowel] != 0 {
		return (*memo)[i][vowel]
	}
	if i == 0 {
		(*memo)[i][vowel] = 1
	} else {
		if vowel == 0 {
			(*memo)[i][vowel] = (_memo(i-1, 1, memo) + _memo(i-1, 2, memo) + _memo(i-1, 4, memo)) % mod
		} else if vowel == 1 {
			(*memo)[i][vowel] = (_memo(i-1, 0, memo) + _memo(i-1, 2, memo)) % mod
		} else if vowel == 2 {
			(*memo)[i][vowel] = (_memo(i-1, 1, memo) + _memo(i-1, 3, memo)) % mod
		} else if vowel == 3 {
			(*memo)[i][vowel] = (_memo(i-1, 2, memo)) % mod
		} else {
			(*memo)[i][vowel] = (_memo(i-1, 2, memo) + _memo(i-1, 3, memo)) % mod
		}
	}
	return (*memo)[i][vowel]
}

//----------------------memoization------------------

func main() {
	fmt.Println(count_vowel_memo(5))
}
