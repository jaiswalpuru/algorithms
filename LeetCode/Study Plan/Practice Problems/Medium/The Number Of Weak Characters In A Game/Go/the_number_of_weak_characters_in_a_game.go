package main

import (
	"fmt"
	"sort"
)

func the_number_of_weak_characters_in_a_game(arr [][]int) int {
	sort.Slice(arr, func(i, j int) bool {
		if arr[i][0] != arr[j][0] {
			return arr[i][0] < arr[j][0]
		}
		return arr[i][1] > arr[j][1]
	})
	n := len(arr)
	res := 0
	stack := [][]int{}
	for i := 0; i < n; i++ {
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			if top[0] == arr[i][0] {
				stack = append(stack, arr[i])
				break
			} else {
				if arr[i][1] > top[1] {
					res++
					stack = stack[:len(stack)-1]
				} else {
					stack = append(stack, arr[i])
					break
				}
			}
		}
		if len(stack) == 0 {
			stack = append(stack, arr[i])
		}
	}
	return res
}

func main() {
	fmt.Println(the_number_of_weak_characters_in_a_game([][]int{
		{5, 5}, {6, 3}, {3, 6},
	}))
}
