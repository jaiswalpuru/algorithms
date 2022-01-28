package main

import (
	"fmt"
)

func goodness(str string, n, k int) int {
	score := 0
	for i := 0; i < n/2; i++ {
		if str[i] != str[n-i-1] {
			score++
		}
	}
	if score == k {
		return 0
	} else if score > k {
		return score - k
	} else {
		return k - score
	}
}

func main() {
	var t int
	line := 1
	fmt.Scanf("%d", &t)
	for t > 0 {
		var (
			n, k int
			str  string
		)
		fmt.Scanf("%d %d", &n, &k)
		fmt.Scanf("%s", &str)
		result := goodness(str, n, k)
		fmt.Printf("Case #%d: %d\n", line, result)
		t--
	}
}
