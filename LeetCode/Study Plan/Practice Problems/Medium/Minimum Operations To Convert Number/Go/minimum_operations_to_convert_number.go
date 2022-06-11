package main

import "fmt"

func minimum_operations_to_convert_number(arr []int, start, goal int) int {
	q := []int{start}
	visited := make(map[int]bool)
	lvl := 0
	for len(q) > 0 {
		temp := []int{}
		for i := 0; i < len(q); i++ {
			if q[i] == goal {
				return lvl
			}
			if q[i] < 0 || q[i] > 1000 {
				continue
			}
			if _, ok := visited[q[i]]; ok {
				continue
			}
			visited[q[i]] = true
			for j := 0; j < len(arr); j++ {
				temp = append(temp, q[i]+arr[j])
				temp = append(temp, q[i]-arr[j])
				temp = append(temp, q[i]^arr[j])
			}
		}
        q = temp
		lvl++
	}
	return -1
}

func main() {
	fmt.Println(minimum_operations_to_convert_number([]int{2, 4, 12}, 2, 12))
}
