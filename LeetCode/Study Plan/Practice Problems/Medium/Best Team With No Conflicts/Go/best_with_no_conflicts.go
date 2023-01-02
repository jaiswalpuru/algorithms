package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	score,
	age int
}

//----------------------memo----------------------------
func best_team_memo(scores []int, ages []int) int {
	sc := make([]Pair, len(scores))
	for i := 0; i < len(scores); i++ {
		sc[i].age = ages[i]
		sc[i].score = scores[i]
	}
	sort.Slice(sc, func(i, j int) bool {
		if sc[i].age == sc[j].age {
			return sc[i].score < sc[j].score
		}
		return sc[i].age < sc[j].age
	})
	memo := make([][]int, len(scores)+1)
	for i := 0; i < len(scores)+1; i++ {
		memo[i] = make([]int, len(scores)+1)
		for j := 0; j < len(scores)+1; j++ {
			memo[i][j] = -1
		}
	}
	return recur(sc, &memo, 0, -1)
}

func recur(sc []Pair, memo *[][]int, start, prev int) int {
	if start == len(sc) {
		return 0
	}
	if (*memo)[prev+1][start] != -1 {
		return (*memo)[prev+1][start]
	}
	if prev == -1 || sc[prev].score <= sc[start].score {
		(*memo)[prev+1][start] = max(sc[start].score+recur(sc, memo, start+1, start), recur(sc, memo, start+1, prev))
	} else {
		(*memo)[prev+1][start] = recur(sc, memo, start+1, prev)
	}
	return (*memo)[prev+1][start]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//----------------------memo----------------------------

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

func main() {
	fmt.Println(best_team_memo([]int{209, 41, 622, 584, 596, 781, 13, 953, 828, 115, 966, 969, 692, 64, 678, 349, 474, 314, 638, 224, 834, 750, 420, 428, 592, 16, 101, 346, 753, 413},
		[]int{99, 69, 81, 58, 40, 44, 29, 46, 53, 87, 82, 91, 30, 37, 42, 24, 59, 86, 17, 83, 70, 5, 55, 54, 23, 88, 88, 44, 29, 17}))
}
