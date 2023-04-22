/*
Design a data structure that accepts a stream of integers and checks if it has a pair of integers that sum up to a particular value.

Implement the TwoSum class:

TwoSum() Initializes the TwoSum object, with an empty array initially.
void add(int number) Adds number to the data structure.
boolean find(int value) Returns true if there exists any pair of numbers whose sum is equal to value, otherwise, it returns false.
*/

package main

import "fmt"

type TwoSum struct {
	arr []int
	mp  map[int]int
}

func Constructor() TwoSum { return TwoSum{arr: make([]int, 0), mp: make(map[int]int)} }

//append in the list
func (this *TwoSum) Add(number int) {
	this.arr = append(this.arr, number)
	this.mp[number] = len(this.arr) - 1
}

//use two sum logic to find
func (this *TwoSum) Find(value int) bool {
	n := len(this.arr)
	for i := 0; i < n; i++ {
		target := value - this.arr[i]
		if _, ok := this.mp[target]; ok && this.mp[target] != i {
			return true
		}
	}
	return false
}

func main() {
	obj := Constructor()
	obj.Add(3)
	obj.Add(2)
	obj.Add(1)
	fmt.Println(obj.Find(2))
	fmt.Println(obj.Find(3))
	fmt.Println(obj.Find(4))
	fmt.Println(obj.Find(5))
	fmt.Println(obj.Find(6))
}
