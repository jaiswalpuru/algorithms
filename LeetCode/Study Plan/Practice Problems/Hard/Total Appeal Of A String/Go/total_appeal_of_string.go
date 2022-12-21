package main

import "fmt"

//--------------------method 1-----------------
func total_appeal_of_string(s string) int64 {
	prev := make([]int, 26)
	res := int64(0)
	for i := 0; i < len(s); i++ {
		prev[int(s[i])-int('a')] = i + 1
		for j := 0; j < 26; j++ {
			res += int64(prev[j])
		}
	}
	return res
}

//--------------------method 1-----------------

//--------------------method 2-----------------
func total_appeal_of_string_eff(s string) int64 {
	n := len(s)
	prev := make([]int, 26)
	total, res := int64(0), int64(0)
	for i := 0; i < n; i++ {
		total += int64(i+1) - int64(prev[int(s[i])-int('a')])
		prev[int(s[i])-int('a')] = i + 1
		res += total
	}
	return res
}

//--------------------method 2-----------------

//--------------------method 3-----------------
func total_appeal_of_string_eff_2(s string) int64 {
	prev := make([]int, 26)
	for i := 0; i < 26; i++ {
		prev[i] = -1
	}
	res := int64(0)
	n := len(s)
	for i := 0; i < n; i++ {
		res += int64(i-prev[int(s[i])-int('a')]) * int64(n-i)
		prev[int(s[i])-int('a')] = i
	}
	return res
}

//--------------------method 3-----------------

func main() {
	fmt.Println(total_appeal_of_string_eff_2("abbca"))
}
