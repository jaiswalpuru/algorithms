package main

import "fmt"

func permutations_with_dups(s string) []string {
	sb := []byte(s)
	res := []string{}
	recur(sb, &res, 0)
	return res
}

func recur(sb []byte, res *[]string, ind int) {
	if ind == len(sb) {
		*res = append(*res, string(sb))
		return
	}
	visited := make(map[byte]bool)
	for i := ind; i < len(sb); i++ {
		if visited[sb[i]] {
			continue
		}
		sb[i], sb[ind] = sb[ind], sb[i]
		recur(sb, res, ind+1)
		sb[i], sb[ind] = sb[ind], sb[i]
		visited[sb[i]] = true
	}
}

func main() {
	fmt.Println(permutations_with_dups("abaac"))
}
