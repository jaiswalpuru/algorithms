package main

import "fmt"

func distribute_candies(candy_type []int) int {
	mp := make(map[int]int)
	half := len(candy_type) / 2
	for i := 0; i < len(candy_type); i++ {
		mp[candy_type[i]]++
	}
	if half < len(mp) {
		return half
	}
	return len(mp)
}

func main() {
	fmt.Println(distribute_candies([]int{1, 1, 2, 2, 3, 3}))
}
