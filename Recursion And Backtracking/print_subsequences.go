//print subsequences

package main

import "fmt"

func print_subsequences(arr []int) {
	_print(arr, []int{}, 0)
}

func _print(arr []int, sub []int, start int) {
	if start >= len(arr) {
		fmt.Println(sub)
		return
	}
	temp := append(sub, arr[start]) //pick the current element
	_print(arr, temp, start+1)
	temp = temp[:len(temp)-1]
	_print(arr, temp, start+1) //remove the current element
}

func main() {
	print_subsequences([]int{3, 1, 2})
}
