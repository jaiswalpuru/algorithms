package main

import "fmt"

func abs(a int) int{
	if a>0{
		return a
	}
	return -1*a
}

func broken_calculator(start, end int) int {
	op := 0
	visited := make(map[int]bool)
	for start != end {
		v1,v2 := start-1, start*2
		op++
		if !visited[abs(v1-end)] && abs(v1-end) < abs(v2-end){
			visited[abs(v1-end)] = true
			start = v1
			
		}else if !visited[abs(v2-end)] && abs(v1-end) > abs(v2-end) {
			visited[abs(v2-end)] = true
			start = v2
		}else if v1 == end || v2 == end {
			break
		}

	}
	return op
}

func main() {
	fmt.Println(broken_calculator(2,3))
}