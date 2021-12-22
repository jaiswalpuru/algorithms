/*
In a deck of cards, each card has an integer written on it.

Return true if and only if you can choose X >= 2 such that it is possible to split the entire deck into 1 or more groups of
cards, where:

Each group has exactly X cards.
All the cards in each group have the same integer.
*/

package main

import "fmt"

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

func deck_of_cards(arr []int) bool {
	n := len(arr)
	mp := make(map[int]int)
	for i := 0; i < n; i++ {
		mp[arr[i]]++
	}

	x := -1
	for _, v := range mp {
		if x == -1 {
			x = v
		} else {
			x = gcd(x, v)
		}
	}

	return x >= 2
}

func main() {
	fmt.Println(deck_of_cards([]int{1, 2, 3, 4, 4, 3, 2, 1}))
	fmt.Println(deck_of_cards([]int{1, 1, 1, 2, 2, 2, 3, 3}))
	fmt.Println(deck_of_cards([]int{1, 1, 2, 2, 2, 2}))
}
