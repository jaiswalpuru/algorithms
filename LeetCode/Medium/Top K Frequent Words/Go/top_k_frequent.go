package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	v   string
	cnt int
}

type P []Pair

func (p P) Len() int { return len(p) }
func (p P) Less(i, j int) bool {
	if p[i].cnt != p[j].cnt {
		return p[i].cnt > p[j].cnt
	}
	return p[i].v < p[j].v
}
func (p P) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func top_k_words(str []string, k int) []string {
	temp := make(P, 0)
	res := []string{}
	n := len(str)
	visited := make(map[string]int)
	for i := 0; i < n; i++ {
		visited[str[i]]++
	}

	for key, val := range visited {
		temp = append(temp, Pair{v: key, cnt: val})
	}
	sort.Sort(temp)
	for i := 0; i < k; i++ {
		res = append(res, temp[i].v)
	}
	return res
}

func main() {
	fmt.Println(top_k_words([]string{
		"i", "love", "leetcode", "i", "love", "coding",
	}, 2))
}
