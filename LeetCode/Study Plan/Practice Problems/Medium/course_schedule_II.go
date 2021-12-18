/*
There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1.
You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if
you want to take course ai.

For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
Return the ordering of courses you should take to finish all courses. If there are many valid answers, return any of them.
If it is impossible to finish all courses, return an empty array.
*/

package main

import "fmt"

func course_schedule(num int, course [][]int) []int {
	mp := make(map[int]int)
	n := len(course)
	for i := 0; i < n; i++ {
		mp[course[i][0]]++
	}

	pre_req := make(map[int][]int)
	for i := 0; i < n; i++ {
		pre_req[course[i][1]] = append(pre_req[course[i][1]], course[i][0])
	}

	indegree := []int{}
	for i := 0; i < num; i++ {
		if _, ok := mp[i]; !ok {
			indegree = append(indegree, i)
		}
	}

	res := []int{}
	for len(indegree) > 0 {
		val := indegree[0]
		res = append(res, val)
		indegree = indegree[1:]
		for i := 0; i < len(pre_req[val]); i++ {
			mp[pre_req[val][i]]--
			if mp[pre_req[val][i]] == 0 {
				indegree = append(indegree, pre_req[val][i])
			}
		}
	}

	for i := 0; i < len(mp); i++ {
		if mp[i] != 0 {
			return []int{}
		}
	}

	return res
}

func main() {
	fmt.Println(course_schedule(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
}
