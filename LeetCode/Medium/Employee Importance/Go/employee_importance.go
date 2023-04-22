package main

import "fmt"

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func employee_importance(e []*Employee, id int) int {
	n := len(e)
	hm_importance := make(map[int]int)
	hm_sub := make(map[int][]int)
	for i := 0; i < n; i++ {
		hm_importance[e[i].Id] = e[i].Importance
		hm_sub[e[i].Id] = append(hm_sub[e[i].Id], e[i].Subordinates...)
	}

	q := []int{id}
	sum := 0
	for len(q) > 0 {
		m := len(q)
		for i := 0; i < m; i++ {
			curr := q[0]
			q = q[1:]
			sum += hm_importance[curr]
			for j := 0; j < len(hm_sub[curr]); j++ {
				q = append(q, hm_sub[curr][j])
			}
		}
	}
	return sum
}

func main() {
	e := []*Employee{
		{1, 5, []int{2, 3}}, {2, 3, []int{}}, {3, 3, []int{}},
	}
	fmt.Println(employee_importance(e, 1))
}
