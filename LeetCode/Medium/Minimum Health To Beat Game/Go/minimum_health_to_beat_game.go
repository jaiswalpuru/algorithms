package main

import "fmt"

func minimum_health_to_beat_game(damage []int, armor int) int64 {
	res := int64(0)
	if len(damage) == 0 {
		return res
	}
	max_val := damage[0]
	max_ind := 0
	for i := 1; i < len(damage); i++ {
		if damage[i] > max_val {
			max_ind = i
			max_val = damage[i]
		}
	}
	if armor >= max_val {
		damage[max_ind] = 0
	} else {
		damage[max_ind] = damage[max_ind] - armor
	}
	for i := 0; i < len(damage); i++ {
		res += int64(damage[i])
	}
	return res + 1
}

func main() {
	fmt.Println(minimum_health_to_beat_game([]int{2, 7, 4, 3}, 4))
}
