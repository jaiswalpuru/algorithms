/*
Given an array of strings names of size n. You will create n folders in your file system such that, at the ith minute,
you will create a folder with the name names[i].

Since two files cannot have the same name, if you enter a folder name that was previously used, the system will have
a suffix addition to its name in the form of (k), where, k is the smallest positive integer such that the obtained name
remains unique.

Return an array of strings of length n where ans[i] is the actual name the system will assign to the ith folder when you
create it.
*/

package main

import (
	"fmt"
)

func make_file_name_unique(arr []string) []string {
	n := len(arr)
	visited := make(map[string]int)
	res := []string{}

	for i := 0; i < n; i++ {
		if _, ok := visited[arr[i]]; ok {
			k := 0
			temp := ""
			for {
				k = visited[arr[i]]
				visited[arr[i]]++

				temp = fmt.Sprintf("%s(%d)", arr[i], k)
				if visited[temp] == 0 {
					break
				}
			}
			visited[temp]++
			res = append(res, temp)
		} else {
			visited[arr[i]]++
			res = append(res, arr[i])
		}
	}
	return res
}

func main() {
	fmt.Println(make_file_name_unique([]string{"gta", "gta(1)", "gta", "avalon"}))
}
