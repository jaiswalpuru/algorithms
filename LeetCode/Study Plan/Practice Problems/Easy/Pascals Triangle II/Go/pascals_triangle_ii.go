package main

import "fmt"

//This wont work for large n
func pascals_traingle(n int) []int {
	res := []int{}
	for i := 0; i <= n; i++ {
		res = append(res, _pascals_value(n, i))
	}
	return res
}

func _pascals_value(row_index, col_index int) int {
	if row_index == 0 || col_index == 0 || row_index == col_index {
		return 1
	}
	return _pascals_value(row_index-1, col_index-1) + _pascals_value(row_index-1, col_index)
}

//More optimized solution
func pascals_traingle_opt(n int) []int {
	res := [][]int{}
	res = append(res, []int{1})
	res = append(res, []int{1, 1})
	return getRowEle(2, n, &res)
}

func getRowEle(ind int, rowIndex int, res *[][]int) []int {
	if rowIndex == 0 {
		return (*res)[0]
	}
	if rowIndex == 1 {
		return (*res)[1]
	}

	temp := []int{1}
	j := 0
	for i := 1; i < ind; i++ {
		temp = append(temp, (*res)[ind-1][j]+(*res)[ind-1][j+1])
		j++
	}
	temp = append(temp, 1)
	*res = append(*res, temp)
	if rowIndex == ind {
		return temp
	}
	return getRowEle(ind+1, rowIndex, res)
}

func main() {
	fmt.Println(pascals_traingle(3))
}
