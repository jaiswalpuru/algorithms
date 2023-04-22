/*
Given a list of non-negative integers nums, arrange them such that they form the largest number and return it.

Since the result may be very large, so you need to return a string instead of an integer.
*/

package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Item []string

func (a Item) Len() int { return len(a) }
func (a Item) Less(i, j int) bool {
	str_one := a[i] + a[j]
	str_two := a[j] + a[i]
	return str_one > str_two
}
func (a Item) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func largest_numebr(arr []int) string {
	n := len(arr)
	str := make(Item, n)
	for i := 0; i < n; i++ {
		str[i] = strconv.Itoa(arr[i])
	}

	sort.Sort(str)
	res := strings.Join(str, "")

	if res[0] == '0' {
		return "0"
	}
	return res
}

func main() {
	fmt.Println(largest_numebr([]int{10, 2}))
	fmt.Println(largest_numebr([]int{3, 30, 34, 5, 9}))
}
