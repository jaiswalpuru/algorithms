/*
You are visiting a farm that has a single row of fruit trees arranged from left to right.
The trees are represented by an integer array fruits where fruits[i] is the type of fruit the ith tree produces.

You want to collect as much fruit as possible. However, the owner has some strict rules that you must follow:

You only have two baskets, and each basket can only hold a single type of fruit.
There is no limit on the amount of fruit each basket can hold.
Starting from any tree of your choice, you must pick exactly one fruit from every
tree (including the start tree) while moving to the right. The picked fruits must fit in one of your baskets.
Once you reach a tree with fruit that cannot fit in your baskets, you must stop.
Given the integer array fruits, return the maximum number of fruits you can pick.

Example 1:
Input: fruits = [1,2,1]
Output: 3
Explanation: We can pick from all 3 trees.

Example 2:
Input: fruits = [0,1,2,2]
Output: 3
Explanation: We can pick from trees [1,2,2].
If we had started at the first tree, we would only pick from trees [0,1].

Example 3:
Input: fruits = [1,2,3,2,2]
Output: 4
Explanation: We can pick from trees [2,3,2,2].
If we had started at the first tree, we would only pick from trees [1,2].

Example 4:
Input: fruits = [3,3,3,1,2,1,1,2,3,3,4]
Output: 5
Explanation: We can pick from trees [1,2,1,1,2].
*/

package main

import "fmt"

func total_fruit(arr []int) int {
	basket := make(map[int]int, 3)
	max_fruits := 0

	n := len(arr)

	for i, j := 0, 0; i < n; i++ {
		basket[arr[i]]++
		for len(basket) > 2 {
			basket[arr[j]]--
			if basket[arr[j]] == 0 {
				delete(basket, arr[j])
			}
			j++
		}
		if i-j+1 > max_fruits {
			max_fruits = i - j + 1
		}
	}
	return max_fruits
}

func main() {
	fmt.Println(total_fruit([]int{1, 2, 1}))
}
