package main

import "fmt"

func findJudge(n int, trust [][]int) int {
	outdegree := make([]int, n+1)
	indegree := make([]int, n+1)
	for i := 0; i < len(trust); i++ {
		outdegree[trust[i][0]]++
		indegree[trust[i][1]]++
	}
	for i := 1; i <= n; i++ {
		if outdegree[i] == 0 && indegree[i] == n-1 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(findJudge(2, [][]int{{1, 2}}))
}
