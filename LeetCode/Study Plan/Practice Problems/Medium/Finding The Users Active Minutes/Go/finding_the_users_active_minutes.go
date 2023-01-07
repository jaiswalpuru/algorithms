package main

import "fmt"

func finding_the_users_active_minutes(logs [][]int, k int) []int {
	uam := make([]int, k)
	logs_group := make(map[int]map[int]int)
	for i := 0; i < len(logs); i++ {
		if _, ok := logs_group[logs[i][0]]; !ok {
			logs_group[logs[i][0]] = make(map[int]int)
		}
		logs_group[logs[i][0]][logs[i][1]]++
	}
	for _, v := range logs_group {
		uam[len(v)-1]++
	}
	return uam
}

func main() {
	fmt.Println(finding_the_users_active_minutes([][]int{
		{0, 5}, {1, 2}, {0, 2}, {0, 5}, {1, 3},
	}, 5))
}
