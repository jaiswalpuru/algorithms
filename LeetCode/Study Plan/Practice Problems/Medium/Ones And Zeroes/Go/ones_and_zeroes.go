package main

import "fmt"

func count_ones_zeroes(s string) []int {
	res := make([]int, 2)
	n := len(s)
	for i := 0; i < n; i++ {
		res[s[i]-'0']++
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//------------------------------brute force------------------------------(TLE)
func ones_and_zeroes(str []string, m, n int) int {
	cnt := 0
	_count_ones_zeroes(str, &cnt, []string{}, 0, m, n)
	return cnt
}

func _count_ones_zeroes(str []string, cnt *int, set []string, ind int, m, n int) {
	if ind == len(str) {
		z, o := 0, 0
		for i := 0; i < len(set); i++ {
			res := count_ones_zeroes(set[i])
			z += res[0]
			o += res[1]
		}
		if z > 0 && o > 0 && z <= m && o <= n {
			*cnt = max(*cnt, len(set))
		}
		return
	}

	temp := append(set, str[ind])
	_count_ones_zeroes(str, cnt, temp, ind+1, m, n)
	temp = temp[:len(temp)-1]
	_count_ones_zeroes(str, cnt, set, ind+1, m, n)
}
//------------------------------brute force------------------------------

//------------------------------memo------------------------------
func ones_and_zeroes_memo(str []string, m,n int) int {
	memo := make([][][]int, len(str))
	for i:=0;i<len(str);i++{
		memo[i] = make([][]int, m+1)
		for j:=0;j<=m;j++{
			memo[i][j] = make([]int, n+1)
		}
	}

	return _count_ones_zeroes_memo(str, 0, m, n, &memo)
}

func _count_ones_zeroes_memo(str []string, ind int, m, n int, memo *[][][]int) int {
	if ind == len(str) {
		return 0
	}
	if (*memo)[ind][m][n] != 0 {
		return (*memo)[ind][m][n]
	}

	res := count_ones_zeroes(str[ind])
	take := -1
	if m-res[0] >=0 && n-res[1] >= 0{
		take = 1 + _count_ones_zeroes_memo(str, ind+1, m-res[0], n-res[1], memo)
	}
	dnt_take := _count_ones_zeroes_memo(str, ind+1, m, n, memo)
	(*memo)[ind][m][n] = max(take, dnt_take)
	return (*memo)[ind][m][n]
}

//------------------------------memo------------------------------

func main() {
	fmt.Println(ones_and_zeroes_memo([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
}
