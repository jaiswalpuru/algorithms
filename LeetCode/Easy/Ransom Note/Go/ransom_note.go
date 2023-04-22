package main

import "fmt"

func ransom_note(a, b string) bool {
	mp_a := make(map[byte]int)
	mp_b := make(map[byte]int)
	n, m := len(a), len(b)
	for i := 0; i < n; i++ {
		mp_a[a[i]]++
	}
	for j := 0; j < m; j++ {
		mp_b[b[j]]++
	}

	for k, v := range mp_a {
		if v > mp_b[k] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(ransom_note("a", "b"))
}
