package main

import "fmt"
import "math"

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minimum_number_of_work_sessions_to_finish_the_tasks(tasks []int, session_time int) int {
	res := math.MaxInt64
	_recur(tasks, session_time, 0, []int{}, &res)
	return res
}

func _recur(tasks []int, session_time, ind int, sess []int, res *int) {
	if ind >= len(tasks) {
		*res = min(*res, len(sess))
		return
	}

	if len(sess) > *res {
		return
	}

	for i := 0; i < len(sess); i++ {
		if sess[i]+tasks[ind] <= session_time {
			sess[i] += tasks[ind]
			_recur(tasks, session_time, ind+1, sess, res)
			sess[i] -= tasks[ind]
		}
	}

	temp := append(sess, tasks[ind])
	_recur(tasks, session_time, ind+1, temp, res)
	temp = temp[:len(temp)-1]
}

func main() {
	fmt.Println(minimum_number_of_work_sessions_to_finish_the_tasks([]int{1, 2, 3}, 3))
}
