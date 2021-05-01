// Given an array of integers and a number k, where 1 <= k <= length of the array,
// compute the maximum values of each subarray of length k.

// For example, given array = [10, 5, 2, 7, 8, 7] and k = 3, we should get: [10, 7, 8, 8], since:

// 10 = max(10, 5, 2)
// 7 = max(5, 2, 7)
// 8 = max(2, 7, 8)
// 8 = max(7, 8, 7)
// Do this in O(n) time and O(k) space.
// You can modify the input array in-place and you do not need to store the results.
// You can simply print them out as you compute them.

/*
 */

package main

import "fmt"

func main() {

	arr := []int{10, 5, 2, 7, 8, 7}
	k := 3

	i := 0
	dequeue := []int{}

	for i = 0; i < k; i++ {
		for len(dequeue) > 0 && arr[i] > arr[dequeue[len(dequeue)-1]] {
			dequeue = dequeue[:len(dequeue)-1]
		}
		dequeue = append(dequeue, i)
	}

	for ; i < len(arr); i++ {
		fmt.Println(arr[dequeue[0]])
		for len(dequeue) > 0 && dequeue[0] <= i-k {
			dequeue = dequeue[1:]
		}

		for len(dequeue) > 0 && arr[i] > arr[dequeue[len(dequeue)-1]] {
			dequeue = dequeue[:len(dequeue)-1]
		}

		dequeue = append(dequeue, i)
	}
	fmt.Println(arr[dequeue[0]])
}
