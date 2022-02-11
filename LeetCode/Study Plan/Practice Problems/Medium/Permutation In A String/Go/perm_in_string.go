package main

import (
	"fmt"
)

//-------------------bad brute force dont't use-------------------
func permutation(s1, s2 string) bool {
	mp := make(map[byte]int)

	n, m := len(s1), len(s2)
	for i := 0; i < n; i++ {
		mp[s1[i]]++
	}

	for i := 0; i < m; i++ {
		if _, ok := mp[s2[i]]; ok {
			visited := make(map[byte]int)
			for j := i; j < m; j++ {
				if _, ok := mp[s2[j]]; ok {
					if visited[s2[j]] > mp[s2[j]] {
						break
					}
					visited[s2[j]]++
				} else {
					break
				}
			}
			fmt.Println(mp, visited)
			f := true
			for k := range mp {
				if mp[k] != visited[k] {
					f = false
					break
				}
			}
			if f {
				return true
			}
		}
	}
	return false
}

//---------------------------------------------------------------------

//------------------efficient approach adopted from the above approach--------------------

func match(s1, s2 []int) bool {
	for i := 0; i < 26; i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func permutation_eff(s1, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}

	s1_map := make([]int, 26)
	for i := 0; i < n; i++ {
		s1_map[s1[i]-'a']++
	}

	for i := 0; i < m-n; i++ {
		s2_map := make([]int, 26)
		for j := 0; j < n; j++ {
			s2_map[s2[i+j]-'a']++
		}
		if match(s1_map, s2_map) {
			return true
		}
	}
	return false
}

//---------------------------------------------------------

func main() {
	fmt.Println(permutation_eff("abc", "cccccbabbbaaaa"))
}
