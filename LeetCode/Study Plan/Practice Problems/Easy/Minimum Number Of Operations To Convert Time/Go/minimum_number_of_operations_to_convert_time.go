package main

import (
	"fmt"
	"strconv"
	"strings"
)

func convert_time(current string, correct string) int {
	curr := strings.Split(current, ":")
	corr := strings.Split(correct, ":")
	hrs1, _ := strconv.Atoi(curr[0])
	hrs2, _ := strconv.Atoi(corr[0])
	min1, _ := strconv.Atoi(curr[1])
	min2, _ := strconv.Atoi(corr[1])

	op := 0
	diff_min := 0
	if min1 > min2 {
		diff_min = (60 - min1) + min2
		hrs1++
	} else {
		diff_min = min2 - min1
	}

	if diff_min >= 15 {
		op += diff_min / 15
		diff_min = (diff_min % 15)
	}

	if diff_min >= 5 {
		op += diff_min / 5
		diff_min = (diff_min % 5)
	}

	if diff_min >= 1 {
		op += diff_min / 1
		diff_min = diff_min % 1
	}

	if hrs2 != hrs1 {
		op += (hrs2 - hrs1)
	}

	return op
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -1 * a
}

func main() {
	fmt.Println(convert_time("02:30", "04:35"))
}
