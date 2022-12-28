package main

import (
	"fmt"
	"sort"
)

type T struct {
	ind int
	src,
	target string
}

func find_and_replace_in_string(s string, indices []int, sources []string, targets []string) string {
	q := make([]T, len(indices))
	for i := 0; i < len(indices); i++ {
		q[i].ind = indices[i]
		q[i].src = sources[i]
		q[i].target = targets[i]
	}
	sort.Slice(q, func(i, j int) bool {
		return q[i].ind < q[j].ind
	})
	ex := 0
	for i := 0; i < len(indices); i++ {
		ind := q[i].ind + ex
		src := q[i].src
		tar := q[i].target
		if ind+len(src) > len(s) {
			continue
		}
		str := s[ind : ind+len(src)]
		if str == src {
			ex += len(tar) - len(src)
			t := s[:ind] + tar + s[ind+len(src):]
			s = t
		}
	}
	return s
}

func main() {
	fmt.Println(find_and_replace_in_string("abcd", []int{0, 2}, []string{"a", "cd"}, []string{"eee", "ffff"}))
}
