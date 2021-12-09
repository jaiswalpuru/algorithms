/*
Given an array arr.  You can choose a set of integers and remove all the occurrences of these integers in the array.

Return the minimum size of the set so that at least half of the integers of the array are removed.

Example 1:

Input: arr = [3,3,3,3,5,5,5,2,2,7]
Output: 2
Explanation: Choosing {3,7} will make the new array [5,5,5,2,2] which has size 5
(i.e equal to half of the size of the old array).
Possible sets of size 2 are {3,5},{3,2},{5,2}.
Choosing set {2,7} is not possible as it will make the new array [3,3,3,3,5,5,5] which has size
greater than half of the size of the old array.


Example 2:

Input: arr = [7,7,7,7,7,7]
Output: 1
Explanation: The only possible set you can choose is {7}. This will make the new array empty.


Example 3:

Input: arr = [1,9]
Output: 1


Example 4:

Input: arr = [1000,1000,3,7]
Output: 1


Example 5:

Input: arr = [1,2,3,4,5,6,7,8,9,10]
Output: 5
*/

package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	num, freq int
}

type P []Pair

func (p P) Len() int           { return len(p) }
func (p P) Less(i, j int) bool { return p[i].freq > p[j].freq }
func (p P) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func reduce_array_size_to_half(arr []int) int {
	res := 0
	mp := make(map[int]int)
	n := len(arr)

	for i := 0; i < n; i++ {
		mp[arr[i]]++
	}

	data := P{}
	for k, v := range mp {
		data = append(data, Pair{num: k, freq: v})
	}

	sort.Sort(data)
	m := len(data)
	freq := data[0].freq
	res = 1
	for i := 1; i < m; i++ {
		if freq >= n/2 {
			return res
		} else {
			freq += data[i].freq
			res++
		}
	}

	return res
}

func main() {
	fmt.Println(reduce_array_size_to_half([]int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7}))
	fmt.Println(reduce_array_size_to_half([]int{7, 7, 7, 7, 7, 7}))
	fmt.Println(reduce_array_size_to_half([]int{1, 9}))
	fmt.Println(reduce_array_size_to_half([]int{1000, 1000, 3, 7}))
	fmt.Println(reduce_array_size_to_half([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}
