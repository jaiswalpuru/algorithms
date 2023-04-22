package main

import (
	"fmt"
	"sort"
)

type Pair struct{ score, age int }

func bestTeamScore(scores []int, ages []int) int {
	size := len(scores)
	scoreAge := make([]Pair, size)

	for i := 0; i < size; i++ {
		scoreAge[i].score = scores[i]
		scoreAge[i].age = ages[i]
	}

	sort.Slice(scoreAge, func(i, j int) bool {
		if scoreAge[i].age == scoreAge[j].age {
			return scoreAge[i].score < scoreAge[j].score
		}
		return scoreAge[i].age < scoreAge[j].age
	})

	memo := make([][]int, len(scoreAge)+1)
	arrayFill(&memo, -1)
	return recur(scoreAge, 0, -1, &memo)
}

func recur(scoreAge []Pair, currIndex, prevIndex int, memo *[][]int) int {
	if currIndex == len(scoreAge) {
		return 0
	}

	if (*memo)[prevIndex+1][currIndex] != -1 {
		return (*memo)[prevIndex+1][currIndex]
	}

	if prevIndex == -1 || scoreAge[prevIndex].score <= scoreAge[currIndex].score {
		(*memo)[prevIndex+1][currIndex] = max(
			scoreAge[currIndex].score+recur(scoreAge, currIndex+1, currIndex, memo),
			recur(scoreAge, currIndex+1, prevIndex, memo))
	} else {
		(*memo)[prevIndex+1][currIndex] = recur(scoreAge, currIndex+1, prevIndex, memo)
	}
	return (*memo)[prevIndex+1][currIndex]
}

func arrayFill(arr *[][]int, value int) {
	for i := 0; i < len(*arr); i++ {
		(*arr)[i] = make([]int, len(*arr))
		for j := 0; j < len(*arr); j++ {
			(*arr)[i][j] = -1
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(bestTeamScore([]int{209, 41, 622, 584, 596, 781, 13, 953, 828, 115, 966, 969, 692, 64, 678, 349, 474, 314, 638, 224, 834, 750, 420, 428, 592, 16, 101, 346, 753, 413},
		[]int{99, 69, 81, 58, 40, 44, 29, 46, 53, 87, 82, 91, 30, 37, 42, 24, 59, 86, 17, 83, 70, 5, 55, 54, 23, 88, 88, 44, 29, 17}))
}
