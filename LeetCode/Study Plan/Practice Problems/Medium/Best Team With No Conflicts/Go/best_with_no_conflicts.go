package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	score,
	age int
}

func best_team(scores []int, ages []int) int {
	n := len(ages)
	score_age := make([]Pair, n)
	for i := 0; i < n; i++ {
		score_age[i].age = ages[i]
		score_age[i].score = scores[i]
	}
	sort.Slice(score_age, func(i, j int) bool {
		if score_age[i].age == score_age[j].age {
			return score_age[i].score < score_age[j].score
		}
		return score_age[i].age < score_age[j].age
	})
	ans := 0
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = score_age[i].score

		for j := 0; j < i; j++ {
			if score_age[i].score >= score_age[j].score {
				dp[i] = max(dp[i], dp[j]+score_age[i].score)
			}
		}
		ans = max(ans, dp[i])
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(best_team([]int{1, 3, 7, 3, 2, 4, 10, 7, 5},
		[]int{4, 5, 2, 1, 1, 2, 4, 1, 4}))
}
