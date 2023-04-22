package main

import "fmt"

//------------worst brute force dont do------------
func remove_dup(arr []int) int {
	n := len(arr)
	mp := make(map[int]int)
	mp_ind := make(map[int]int)
	for i := 0; i < n; i++ {
		if _, ok := mp[arr[i]]; !ok {
			mp_ind[i] = arr[i]
		}
		mp[arr[i]]++
	}

	res := []int{}

	for i := 0; i < n; i++ {
		if _, ok := mp_ind[i]; ok {
			val, cnt := mp_ind[i], mp[arr[i]]
			fmt.Println(val, cnt)
			if cnt > 2 {
				cnt = 2
			}
			for j := 0; j < cnt; j++ {
				res = append(res, val)
			}
		}
	}
	for i := 0; i < len(res); i++ {
		arr[i] = res[i]
	}
	return len(res)
}

//------------------------------------------------

//----------------without using any extra space-----------------------
func remove_dup_eff(arr []int) int {
	n := len(arr)
	j, cnt := 1, 1

	for i := 1; i < n; i++ {
		if arr[i] == arr[i-1] {
			cnt++
		} else {
			cnt = 1
		}
		if cnt <= 2 {
			arr[j] = arr[i]
			j++
		}
	}
	return j
}

//---------------------------------------------------------------------

func main() {
	fmt.Println(remove_dup([]int{1, 1, 1, 2, 2, 3}))
}
