package main

import "fmt"

func kth_lexicographical_string(n, k int) string {
	str := []string{"a", "b", "c"}
	res := []string{}
	_kth_lex_string(n, &res, str, "")
	if len(res) < k {
		return ""
	}
	return res[k-1]
}

func _kth_lex_string(n int, res *[]string, str []string, s string) {
	if len(s) == n {
		*res = append(*res, s)
		return
	}

	for i := 0; i < len(str); i++ {
		if len(s) > 0 && string(s[len(s)-1]) == str[i] {
			continue
		}
		temp := s + str[i]
		_kth_lex_string(n, res, str, temp)
		temp = temp[:len(temp)-1]
	}
}

func main() {
	fmt.Println(kth_lexicographical_string(3, 9))
}
