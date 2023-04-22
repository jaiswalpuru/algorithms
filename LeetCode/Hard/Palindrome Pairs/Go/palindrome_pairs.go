package main

import "fmt"

//-----------------brute force-----------------------
func palindrome_pairs(words []string) [][]int {
	n := len(words)
	res := [][]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j {
				str := words[i] + words[j]
				if check_palindrome(str) {
					res = append(res, []int{i, j})
				}
			}
		}
	}
	return res
}

func check_palindrome(word string) bool {
	i, j := 0, len(word)-1
	for i < j {
		if word[i] != word[j] {
			return false
		}
		i++
		j--
	}
	return true
}

//-----------------brute force-----------------------

//-------------------a better approach-----------------------
func palindrome_pairs_better(words []string) [][]int {
	res := [][]int{}
	mp := make(map[string]int)
	n := len(words)
	for i := 0; i < n; i++ {
		mp[words[i]] = i
	}
	for k, curr_word_ind := range mp {
		reverse_word := reverse(k)
		if _, ok := mp[reverse_word]; ok {
			if mp[reverse_word] != curr_word_ind {
				res = append(res, []int{curr_word_ind, mp[reverse_word]})
			}
		}
		suffix := all_suffix(k)
		prefix := all_prefix(k)
		for i := 0; i < len(suffix); i++ {
			reverse_suff := reverse(suffix[i])
			if _, ok := mp[reverse_suff]; ok {
				res = append(res, []int{mp[reverse_suff], curr_word_ind})
			}
		}
		for i := 0; i < len(prefix); i++ {
			reverse_pre := reverse(prefix[i])
			if _, ok := mp[reverse_pre]; ok {
				res = append(res, []int{curr_word_ind, mp[reverse_pre]})
			}
		}
	}
	return res
}

func all_prefix(str string) []string {
	res := []string{}
	for i := 0; i < len(str); i++ {
		if _check_palindrome(str, i, len(str)-1) {
			res = append(res, str[0:i])
		}
	}
	return res
}

func all_suffix(str string) []string {
	res := []string{}
	for i := 0; i < len(str); i++ {
		if _check_palindrome(str, 0, i) {
			res = append(res, str[i+1:])
		}
	}
	return res
}

func _check_palindrome(word string, i, j int) bool {
	for i < j {
		if word[i] != word[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func reverse(str string) string {
	b := []byte(str)
	i, j := 0, len(str)-1
	for i < j {
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}
	return string(b)
}

//-------------------a better approach-----------------------

func main() {
	fmt.Println(palindrome_pairs_better([]string{"abcd", "dcba", "lls", "s", "sssll"}))
}
