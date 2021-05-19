/*
	Given an array of strictly the characters 'R', 'G', and 'B', segregate the values of the array
	so that all the Rs come first, the Gs come second, and the Bs come last.
	You can only swap elements of the array.

	Do this in linear time and in-place.

	For example, given the array ['G', 'B', 'R', 'R', 'B', 'R', 'G'],
	it should become ['R', 'R', 'R', 'G', 'G', 'B', 'B'].

	This problem is also known as Dutch national flag problem.
*/

package main

import "fmt"

//lets say if there were only two partition i.e R and G
func partition_two(arr []string) []string {
	low, high := 0, len(arr)-1
	for low <= high {
		if arr[low] == "R" {
			low += 1
		} else {
			arr[low], arr[high] = arr[high], arr[low]
			high -= 1
		}
	}
	return arr
}

//now for partition variation we have three characters
func partition_three(arr []string) []string {
	low, mid, high := 0, 0, len(arr)-1
	for mid <= high {
		if arr[mid] == "R" {
			arr[mid], arr[low] = arr[low], arr[mid]
			low += 1
			mid += 1
		} else if arr[mid] == "G" {
			mid += 1
		} else {
			arr[mid], arr[high] = arr[high], arr[mid]
			high -= 1
		}
	}
	return arr
}

func main() {
	str := []string{"R", "G", "G", "R", "G"}
	fmt.Println(partition_two(str))
	arr := []string{"G", "B", "R", "R", "B", "R", "G"}
	fmt.Println(partition_three(arr))
}
