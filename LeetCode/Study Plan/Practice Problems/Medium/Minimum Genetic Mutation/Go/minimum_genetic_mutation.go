package main

import "fmt"

func minimum_genetic_mutation(start, end string, bank []string) int {
	mp := make(map[string]bool)
	q := []string{start}
	b := make(map[string]bool)
	for i := 0; i < len(bank); i++ {
		b[bank[i]] = true
	}
	mp[start] = true
	arr := []rune{'A', 'C', 'G', 'T'}
	cnt := 0
	for len(q) > 0 {
		n := len(q)
		for k := 0; k < n; k++ {
			curr := q[0]
			q = q[1:]
			if curr == end {
				return cnt
			}
			for j := 0; j < 4; j++ {
				for i := 0; i < len(curr); i++ {
					temp := curr[0:i] + string(arr[j]) + curr[i+1:]
					if !mp[temp] && b[temp] {
						q = append(q, temp)
						mp[temp] = true
					}
				}
			}
		}
		cnt++
	}
	return -1
}

func main() {
	fmt.Println(minimum_genetic_mutation("AACCGGTT", "AACCGGTA", []string{"AACCGGTA"}))
}
