package main

import (
	"fmt"
	"sort"
)

func wiggle_sort(arr []int) {
	temp := make([]int, len(arr))
	copy(temp, arr)
	sort.Ints(temp)
	temp_rev := make([]int, len(arr))
	k:=0
	for i:=len(arr)-1; i>=0; i-- {
		temp_rev[i] = temp[k]
		k++
	}
	j,k :=0, 0
	for i:=0;i<len(arr); i++{
		if i%2==0{
			arr[i] = temp[j]
			j++
		}else{
			arr[i] = temp_rev[k]
			k++
		}
	}
	fmt.Println(arr)
}

func main() {
	wiggle_sort([]int{3,5,2,1,6,4})
}