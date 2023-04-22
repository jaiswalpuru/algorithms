package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func minimum_time_differece(arr []string) int {
	min_diff := math.MaxInt64
	a := []int{}
	for i := 0; i < len(arr); i++ {
		curr := strings.Split(arr[i], ":")
		hr, _ := strconv.Atoi(curr[0])
		min, _ := strconv.Atoi(curr[1])
		a = append(a, hr*60+min)
	}
	sort.Ints(a)
	a = append(a, a[0]+(24*60))
	for i := 1; i < len(a); i++ {
		min_diff = min(min_diff, a[i]-a[i-1])
	}
	return min_diff
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(minimum_time_differece([]string{"23:59", "00:00"}))
}
