package main

import "fmt"

func kill_process(pid, ppid []int, kill int) []int {
	relation := make(map[int][]int)
	n := len(ppid)

	for i := 0; i < n; i++ {
		relation[ppid[i]] = append(relation[ppid[i]], pid[i])
	}

	res := []int{}
	q := []int{kill}
	for len(q) > 0 {
		curr := q[0] //parent
		q = q[1:]
		res = append(res, curr)
		children := relation[curr]
		q = append(q, children...)
	}
	return res
}

func main() {
	fmt.Println(kill_process([]int{1, 3, 10, 5}, []int{3, 0, 5, 3}, 5))
}
