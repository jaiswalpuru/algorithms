package main

import "fmt"

func zig_zag(str string, rows int) string {
	if rows == 1 {
		return str
	}

	res := ""
	n := len(str)
	cycle_len := 2*rows - 2

	for i := 0; i < rows; i++ {
		for j := 0; j+i < n; j += cycle_len {
			res += string(str[j+i])
			if i != 0 && i != rows-1 && j+cycle_len-i < n {
				res += string(str[j+cycle_len-i])
			}
		}
	}
	return res
}

func zig_zag_brute_force(str string, rows int) string {
	n := len(str)
	res := make([][]string, rows)
	if len(str) <= 2 {
		return str
	}
	for i := 0; i < rows; i++ {
		res[i] = make([]string, n)
	}
	i, j, k := 0, 0, 1
	res[0][0] = string(str[0])
	for k < n {
		for i = 1; i < rows && k < n; i++ {
			res[i][j] = string(str[k])
			k++
		}
		j++
		i -= 2
		fmt.Println(j)
		for p := i; p >= 0 && k < n; p-- {
			res[p][j] = string(str[k])
			j++
			k++
		}
	}

	ans := ""
	for i := 0; i < rows; i++ {
		for j := 0; j < n; j++ {
			if res[i][j] != "" {
				ans += res[i][j]
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(zig_zag("PAYPALISHIRING", 4))
}
