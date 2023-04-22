package main

import "fmt"

func count_number_of_distinct_integers_after_reverse_operations(arr []int) int {
	n := len(arr)
	mp := make(map[int]bool)
	for i := 0; i < n; i++ {
		num := arr[i]
		mp[num] = true
		val := reverse(num)
		if !mp[val] {
			mp[val] = true
		}
	}
	return len(mp)
}

func reverse(a int) int {
	rev := 0
	for a > 0 {
		temp := a % 10
		rev = rev*10 + temp
		a = a / 10
	}
	return rev
}

func main() {
	fmt.Println(count_number_of_distinct_integers_after_reverse_operations([]int{1, 13, 10, 12, 31}))
}
