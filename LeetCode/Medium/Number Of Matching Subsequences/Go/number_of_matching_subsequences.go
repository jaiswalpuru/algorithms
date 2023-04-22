package main

import "fmt"

//----------------------------brute force------------------------------------
func number_of_matching_subsequences(str string, arr []string) int {
	mp := make(map[string]int)
	n := len(arr)
	for i := 0; i < n; i++ {
		mp[arr[i]]++
	}
	cnt := 0
	visited := make(map[string]bool)
	_recur(str, &mp, "", 0, &cnt, &visited)
	return cnt
}

func _recur(str string, mp *map[string]int, s string, ind int, cnt *int, visited *map[string]bool) {
	if ind == len(str) {
		if (*mp)[string(s)] > 0 && !(*visited)[string(s)] {
			(*cnt) += (*mp)[string(s)]
			(*visited)[string(s)] = true
		}
		return
	}

	temp := s + string(str[ind])
	_recur(str, mp, temp, ind+1, cnt, visited)
	temp = temp[:len(temp)-1]
	_recur(str, mp, temp, ind+1, cnt, visited)
}

//----------------------------brute force------------------------------------

//----------------------------brute force 2------------------------------------
func number_of_matching_subsequences_bf2(str string, arr []string) int {
	mp := make(map[string]int)
	n := len(arr)
	for i := 0; i < n; i++ {
		mp[arr[i]]++
	}

	cnt := 0
	for k, v := range mp {
		i, j, m := 0, 0, len(str)
		for i < m && j < len(k) {
			if k[j] == str[i] {
				i++
				j++
			} else {
				i++
			}
		}
		if j == len(k) {
			cnt += v
		}
	}
	return cnt
}

//----------------------------brute force 2------------------------------------

func main() {
	fmt.Println(number_of_matching_subsequences_bf2("abcde", []string{"a", "bb", "acd", "ace"}))
}
