package main

import "fmt"

func one_edit_distance(s1, s2 string) bool {
	n, m := len(s1), len(s2)

	if n > m {
		s1, s2 = s2, s1
		n, m = m, n
	}

	if m-n > 1 {
		return false
	}

	if m-n == 1 {
		cnt := 0
		i, j := 0, 0
		for i < n && j < m {
			if s1[i] == s2[j] {
				i++
				j++
			} else if s1[i] != s2[j] {
				j++
				cnt++
			}
		}
		if cnt > 1 {
			return false
		}
		return true
	} else {
		cnt := 0
		i, j := 0, 0
		for i < n && j < m {
			if s1[i] != s2[j] {
				cnt++
				if cnt > 1 {
					return false
				}
			}
			i++
			j++
		}
		return true
	}
}

func main() {
	fmt.Println(one_edit_distance("aca", "caca"))
}
