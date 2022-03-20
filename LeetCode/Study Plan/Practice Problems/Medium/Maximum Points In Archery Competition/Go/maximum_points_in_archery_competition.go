package main

import "fmt"

func max_points_in_archery_competition(num_arrows int, alice []int) []int {
	bob := []int{}
	sum := 0
	temp := make([]int, 12)
	_max_points_in_archery_competition(num_arrows, alice, &bob, temp, 11, &sum, 0)
	return bob
}

func _max_points_in_archery_competition(num_arrows int, alice []int, bob *[]int, temp []int, n int, target *int, sum int){
	if n == -1 || num_arrows == 0 {
		if *target < sum {
			*target = sum
			if num_arrows > 0 {
				temp[0] += num_arrows
			}
			*bob = append([]int{}, temp...)
		}
		return
	}

	take := alice[n]+1
	if take <= num_arrows {
		temp[n] = take
		_max_points_in_archery_competition(num_arrows-take, alice, bob, temp, n-1, target, sum+n)
		temp[n] = 0
	}
	_max_points_in_archery_competition(num_arrows, alice, bob, temp, n-1, target, sum)
}

func main() {
	fmt.Println(max_points_in_archery_competition(9, []int{1,1,0,1,0,0,2,1,0,1,2,0}))
}