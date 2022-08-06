package main

import (
	"fmt"
	"math"
)

func poor_pigs(buckets int, min_to_die int, min_to_test int) int {
	states := min_to_test/min_to_die + 1
	return int(math.Ceil(math.Log(float64(buckets)) / math.Log(float64(states))))
}

func main() {
	fmt.Println(poor_pigs(1000, 15, 60))
}
