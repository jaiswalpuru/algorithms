package main

import "fmt"

func minimum_rounds_to_complete_all_tasks(tasks []int) int {
	freq := make(map[int]int)
	for i := 0; i < len(tasks); i++ {
		freq[tasks[i]]++
	}
	rounds := 0
	for _, v := range freq {
		if v == 1 {
			return -1
		}
		if v%3 == 0 {
			rounds += v / 3
		} else {
			rounds += (v / 3) + 1
		}
	}
	return rounds
}

func main() {
	fmt.Println(minimum_rounds_to_complete_all_tasks([]int{2, 2, 3, 3, 2, 4, 4, 4, 4, 4}))
}
