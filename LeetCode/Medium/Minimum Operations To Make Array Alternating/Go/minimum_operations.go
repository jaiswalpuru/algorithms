package main

import "fmt"

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max_find(mp map[int]int) (int, int, int) {
	f1, f2, max_val := 0, 0, 0
	for k, v := range mp {
		if v >= f1 {
			f2 = f1
			f1 = v
			max_val = k
		} else if v >= f2 {
			f2 = v
		}
	}
	return f1, f2, max_val
}

func minimum_operation(arr []int) int {
	mp_even := make(map[int]int)
	mp_odd := make(map[int]int)
	n := len(arr)
	cnt_even := 0
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			mp_even[arr[i]]++
			cnt_even++
		} else {
			mp_odd[arr[i]]++
		}
	}

	even_one, even_two, max_even := max_find(mp_even)
	odd_one, odd_two, max_odd := max_find(mp_odd)
	fmt.Println(even_one, even_two, odd_one, odd_two, max_even, max_odd)
	even_one = cnt_even - even_one
	even_two = cnt_even - even_two
	odd_one = n - cnt_even - odd_one
	odd_two = n - cnt_even - odd_two

	if max_even != max_odd {
		return even_one + odd_one
	}
	return min(even_one+odd_two, even_two+odd_one)
}

func main() { fmt.Println(minimum_operation([]int{3, 1, 3, 2, 4, 3})) }
