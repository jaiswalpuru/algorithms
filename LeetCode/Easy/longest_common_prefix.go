/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".
*/

package main

import "fmt"

func lcp(arr []string) string {

	shortest_word := arr[0]
	ind := 0
	n := len(arr)

	for i := 1; i < n; i++ {
		if len(shortest_word) < len(arr[i]) {
			shortest_word = arr[i]
			ind = i
		}
	}

	for i := 0; i < n; i++ {
		if i != ind {
			if shortest_word != arr[i] {
				j, k := 0, 0
				n, m := len(shortest_word), len(arr[i])
				temp := ""
				for j < n && k < m {
					if shortest_word[j] != arr[i][j] {
						break
					}
					temp += string(shortest_word[j])
					j++
					k++
				}
				shortest_word = temp
			}
		}
	}
	return shortest_word
}

func main() {
	fmt.Println(lcp([]string{"cir", "car"}))
}
