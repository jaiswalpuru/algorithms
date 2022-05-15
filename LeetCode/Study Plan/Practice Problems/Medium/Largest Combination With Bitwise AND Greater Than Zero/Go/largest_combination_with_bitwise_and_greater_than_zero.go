package main

import "fmt"

func largest_combination(arr []int) int {
	bits := make([]int, 32)
	for i:=0;i<len(arr);i++{
		t := 31
		for arr[i] > 0 {
			bits[t] += arr[i]%2
			arr[i] = arr[i]/2
			t--
		}
	}
	
	res :=0
	for i:=0;i<len(bits);i++{
		res = max(res, bits[i])
	}
	return res
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}

func main(){
	fmt.Println(largest_combination([]int{16,17,71,62,12,24,14}))
}