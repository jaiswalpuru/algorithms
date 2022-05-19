package main

import "fmt"
import "math"

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}

func is_safe(i,j,n,m, parent int, arr [][]int) bool {
	return i>=0 && j>=0 && i<n && j<m && parent < arr[i][j]
}

var dir = [][]int{{0,1},{0,-1}, {1,0}, {-1,0}}


//--------------------------naive recursive approach--------------------------
func dfs_naive(p, q, n,m int, arr [][]int) int {
	path := 0
	for i:=0;i<4;i++{
		x,y := p+dir[i][0], q+dir[i][1]
		if is_safe(x, y, n, m, arr[p][q], arr) {
            path = max(path, dfs_naive(x,y,n,m, arr))
		}
	}
	return path+1
}

func longest_increasing_path_in_matrix_naive(arr [][]int) int {
	n,m := len(arr), len(arr[0])

	res := math.MinInt64
	for i:=0;i<n; i++ {
		for j:=0;j<m;j++{
			res = max(res, dfs_naive(i,j, n,m, arr))
		}
	}
	return res
}
//--------------------------naive recursive approach--------------------------


//--------------------------memoization----------------------------------
func dfs(p, q, n,m int, arr [][]int, memo *[][]int) int {
	
	if (*memo)[p][q] != 0 {
		return (*memo)[p][q]
	}
	
	for i:=0;i<4;i++{
		x,y := p+dir[i][0], q+dir[i][1]
		if is_safe(x, y, n, m, arr[p][q], arr) {
        	(*memo)[p][q] = max((*memo)[p][q], dfs(x,y,n,m, arr, memo))
		}
	}
	(*memo)[p][q] += 1
	return (*memo)[p][q]
}

func longest_increasing_path_in_matrix(arr [][]int) int {
	n,m := len(arr), len(arr[0])
	memo := make([][]int, n)
	
	for i:=0; i<n; i++{
		memo[i] = make([]int, m)
	}

	res := math.MinInt64
	for i:=0;i<n; i++ {
		for j:=0;j<m;j++{
			res = max(res, dfs(i,j, n,m, arr, &memo))
		}
	}
	return res
}
//--------------------------memoization-----------------------------------

func main(){
	fmt.Println(longest_increasing_path_in_matrix([][]int{
		{9,9,4},{6,6,8},{2,1,1},
	}))
}