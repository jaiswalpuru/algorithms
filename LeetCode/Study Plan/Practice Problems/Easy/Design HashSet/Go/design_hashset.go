package main

type MyHashSet struct {
	arr []int
}

func Constructor() MyHashSet {
	arr := make([]int, 1e6+1)
	for i := 0; i <= 1e6; i++ {
		arr[i] = -1
	}
	return MyHashSet{arr: arr}
}

func (this *MyHashSet) Add(key int) { this.arr[key] = key }

func (this *MyHashSet) Remove(key int) { this.arr[key] = -1 }

func (this *MyHashSet) Contains(key int) bool {
	if this.arr[key] != -1 {
		return true
	}
	return false
}

func main() {

}
