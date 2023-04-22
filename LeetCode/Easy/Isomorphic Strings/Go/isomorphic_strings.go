package main

import "fmt"

func isomorphic_strings(s1, s2 string) bool {
	mp := make(map[byte]byte)
	mp_r := make(map[byte]byte)
	for i := 0; i < len(s1); i++ {
		if _, ok := mp[s1[i]]; ok {
			if s2[i] != mp[s1[i]] {
				return false
			}
		}
		if _, ok := mp_r[s2[i]]; ok {
			if s1[i] != mp_r[s2[i]] {
				return false
			}
		}
		mp[s1[i]] = s2[i]
		mp_r[s2[i]] = s1[i]
	}
	return true
}

func main() {
	fmt.Println(isomorphic_strings("egg", "add"))
}
