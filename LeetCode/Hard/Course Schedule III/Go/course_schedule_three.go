package main

import "fmt"
import "sort"

type I [][]int

func (a I) Len() int           { return len(a) }
func (a I) Less(i, j int) bool { return a[i][1] < a[j][1] }
func (a I) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//-------------------brute force-------------------
func course_schedule_three(courses [][]int) int {
	var a I
	a = courses
	sort.Sort(a)
	return _recur(a, 0, 0)
}

func _recur(c [][]int, ind int, time int) int {
	if ind == len(c) {
		return 0
	}
	taken := 0
	if time+c[ind][0] <= c[ind][1] {
		taken = 1 + _recur(c, ind+1, time+c[ind][0])
	}
	not_take := _recur(c, ind+1, time)
	return max(taken, not_take)
}

//-------------------brute force-------------------

//-------------------memo------------------ memory issue
func course_schedule_three_memo(c [][]int) int {
	var a I
	a = c
	sort.Sort(a)
	n := len(c)
	memo := make([][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, c[n-1][1]+1)
		for j := 0; j < c[n-1][1]+1; j++ {
			memo[i][j] = -1
		}
	}

	return _memo(a, 0, 0, &memo)
}

func _memo(a [][]int, ind, time int, memo *[][]int) int {
	if ind == len(a) {
		return 0
	}
	if (*memo)[ind][time] != -1 {
		return (*memo)[ind][time]
	}
	take := 0
	if time+a[ind][0] <= a[ind][1] {
		take = 1 + _memo(a, ind+1, time+a[ind][0], memo)
	}
	not_take := _memo(a, ind+1, time, memo)
	(*memo)[ind][time] = max(take, not_take)
	return (*memo)[ind][time]
}

//-------------------memo------------------

//-----------------------iterative approach------------------
func course_schedule_three_iterative(c [][]int) int {
	var a I
	a = c
	sort.Sort(a)
	n := len(c)
	time, cnt := 0, 0
	for i := 0; i < n; i++ {
		if time+c[i][0] <= c[i][1] {
			time += c[i][0]
			cnt++
		} else {
			max_ind := i
			for j := 0; j < i; j++ {
				if c[j][0] > c[max_ind][0] {
					max_ind = j
				}
			}
			if c[max_ind][0] > c[i][0] {
				time += c[i][0] - c[max_ind][0]
			}
			c[max_ind][0] = -1
		}
	}
	return cnt
}

//-----------------------iterative approach------------------

func main() {
	fmt.Println(course_schedule_three_iterative([][]int{
		{100, 200}, {200, 1300}, {1000, 1250}, {2000, 3200},
	}))
}
