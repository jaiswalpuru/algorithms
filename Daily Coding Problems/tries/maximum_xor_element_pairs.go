package main

import "fmt"

//This question can also be done using trie also

//The approach below is similar to find max and pairs
func main() {
	arr := []int{4, 6, 7}
	fmt.Println(max_and_pair(arr))
	fmt.Println(max_xor_pair(arr))
}

func max_xor_pair(arr []int) int {
	max, mask := 0, 0
	mp := make(map[int]int)

	for i := 30; i >= 0; i-- {

		//set the i'th bit in the mask
		mask |= (1 << i)

		for j := 0; j < len(arr); j++ {
			//keep the prefix till the ith bit neglecting all the bits after that
			mp[arr[j]&mask] = 1
		}

		new_max := max | (1 << i)
		for k := range mp {
			if mp[k^new_max] == 1 {
				max = new_max
				break
			}
		}
		mp = make(map[int]int)
	}
	return max
}

func max_and_pair(arr []int) int {
	res, count := 0, 0
	for bit := 31; bit >= 0; bit-- {
		count = check_bit(res|(1<<bit), arr)
		if count >= 2 {
			res |= (1 << bit)
		}
	}
	return res
}

func check_bit(pattern int, arr []int) int {
	cnt := 0
	for i := 0; i < len(arr); i++ {
		if (pattern & arr[i]) == pattern {
			cnt++
		}
	}
	return cnt
}
