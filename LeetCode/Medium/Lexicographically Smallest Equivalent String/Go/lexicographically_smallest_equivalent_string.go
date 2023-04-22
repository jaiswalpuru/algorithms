package main

import "fmt"

func lexicographically_smallest_equivalent_string(s1, s2, base_str string) string {
	parent := make([]int, 26)
	size := make([]int, 26)
	for i := 0; i < 26; i++ {
		parent[i] = i
		size[i] = 1
	}
	for i := 0; i < len(s1); i++ {
		if find(int(s1[i]-'a'), &parent) != find(int(s2[i]-'a'), &parent) {
			union(int(s1[i]-'a'), int(s2[i]-'a'), &parent, &size)
		}
	}
	mp := make(map[int][]int)
	for i := 0; i < 26; i++ {
		mp[parent[i]] = append(mp[parent[i]], i)
	}
	res := ""
	for i := 0; i < len(base_str); i++ {
		p := find(int(base_str[i]-'a'), &parent)
		v := mp[p]
		if len(v) > 0 {
			res += string(v[0] + 'a')
		} else {
			res += string(base_str[i])
		}
	}
	return res
}

func union(x, y int, parent *[]int, size *[]int) {
	x = find(x, parent)
	y = find(y, parent)
	if x == y {
		return
	}
	if (*size)[x] >= (*size)[y] {
		(*size)[x] += (*size)[y]
		(*parent)[y] = x
	} else {
		(*size)[y] += (*size)[x]
		(*parent)[x] = y
	}
}

func find(x int, parent *[]int) int {
	if (*parent)[x] != x {
		(*parent)[x] = find((*parent)[x], parent)
	}
	return (*parent)[x]
}

func main() {
	fmt.Println(lexicographically_smallest_equivalent_string("parker", "morris", "parser"))
}
