/*
There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1.
You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must
take course bi first if you want to take course ai.

For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
Return true if you can finish all courses. Otherwise, return false.

Example 1:
Input: numCourses = 2, prerequisites = [[1,0]]
Output: true
Explanation: There are a total of 2 courses to take.
To take course 1 you should have finished course 0. So it is possible.

Example 2:
Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
Output: false
Explanation: There are a total of 2 courses to take.
To take course 1 you should have finished course 0, and to take course 0 you should also have finished course 1.
So it is impossible.
*/

package main

import (
	"fmt"
)

func course_schedule(num_courses int, pre_req [][]int) bool {
	n := len(pre_req)
	dependencies := make([]int, num_courses)

	g := make(map[int][]int)
	for i := 0; i < n; i++ {
		dependencies[pre_req[i][0]]++
		if _, ok := g[pre_req[i][1]]; ok {
			g[pre_req[i][1]] = append(g[pre_req[i][1]], pre_req[i][0])
		} else {
			g[pre_req[i][1]] = append([]int{}, pre_req[i][0])
		}
	}

	q := []int{}
	for i := 0; i < num_courses; i++ {
		if dependencies[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		c := q[0]
		sub_class := g[c]
		for i := 0; i < len(sub_class); i++ {
			dependencies[sub_class[i]]--
			if dependencies[sub_class[i]] == 0 {
				q = append(q, sub_class[i])
			}
		}
		q = q[1:]
	}

	for i := 0; i < num_courses; i++ {
		if dependencies[i] != 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(course_schedule(2, [][]int{{1, 0}}))
	fmt.Println(course_schedule(2, [][]int{{1, 0}, {0, 1}}))
	fmt.Println(course_schedule(20, [][]int{{0, 10}, {3, 18}, {5, 5}, {6, 11}, {11, 14}, {13, 1}, {15, 1}, {17, 4}}))
}
