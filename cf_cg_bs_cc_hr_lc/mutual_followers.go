/*
You are given a two-dimensional list of integers relations.
Each element relations[i] contains [a, b] meaning that person a is following person b on Twitter.

Return the list of people who follow someone that follows them back, sorted in ascending order.

Constraints

0 ≤ n ≤ 100,000 where n is the length of relations
Example 1
Input
relations = [
    [0, 1],
    [2, 3],
    [3, 4],
    [1, 0]
]
Output
[0, 1]
Explanation
0 follows 1 and 1 follows 0.
Example 2
Input
relations = [
    [0, 1],
    [1, 2],
    [2, 3],
    [3, 0]
]
Output
[]
Explanation
There aren't any mutual followers.
*/

package main

import (
	"fmt"
	"sort"
)

func search(val int, arr []int) bool {
	for i := 0; i < len(arr); i++ {
		if val == arr[i] {
			return true
		}
	}
	return false
}

func main() {

	relation := [][]int{
		{0, 1},
		{0, 2},
		{1, 0},
		{2, 0},
	}

	mp := make(map[int][]int)
	mutualFollowers := []int{}
	visited := make(map[int]bool)

	for _, y := range relation {
		mp[y[0]] = append(mp[y[0]], y[1])
		if _, ok := mp[y[1]]; ok {
			if search(y[0], mp[y[1]]) {
				if !visited[y[0]] {
					mutualFollowers = append(mutualFollowers, y[0])
					visited[y[0]] = true
				}
				if !visited[y[1]] {
					mutualFollowers = append(mutualFollowers, y[1])
					visited[y[1]] = true
				}
			}
		}
	}
	sort.Ints(mutualFollowers)
	fmt.Println(mutualFollowers)
}
