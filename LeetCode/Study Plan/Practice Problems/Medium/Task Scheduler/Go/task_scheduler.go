package main

import (
	"fmt"
	"sort"
)

func task_scheduler(tasks []byte, n int) int {
	freq := make([]int, 26)
	for i := 0; i < len(tasks); i++ {
		freq[int(tasks[i])-int('A')]++
	}
	sort.Ints(freq)
	max_freq := freq[25]
	idle_time := (max_freq - 1) * n
	for i := len(freq) - 2; i >= 0; i-- {
		idle_time -= min(max_freq-1, freq[i])
	}
	idle_time = max(idle_time, 0)
	return idle_time + len(tasks)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(task_scheduler([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2))
}
