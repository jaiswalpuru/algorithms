package main

import "fmt"

func replace_elements_in_array(arr []int, op [][]int) []int {
	index := make(map[int]int)
	n := len(arr)
	for i := 0; i < n; i++ {
		index[arr[i]] = i
	}
    
	m := len(op)
	for i := 0; i < m; i++ {
		ind := index[op[i][0]]
		index[op[i][1]] = ind
		delete(index, op[i][0])
	}

    res := make([]int, n)
    for k,v := range index {
        res[v] = k
    }
	
    return res
}

func main() {
	fmt.Println(replace_elements_in_array([]int{1, 2, 4, 6}, [][]int{
		{1, 3}, {4, 7}, {6, 1},
	}))
}
