package main

type MyHashMap struct{ arr []int }

func Constructor() MyHashMap {
	arr := make([]int, 1e6+1)
	for i := 0; i <= 1e6; i++ {
		arr[i] = -1
	}
	return MyHashMap{arr: arr}
}

func (this *MyHashMap) Put(key int, value int) { this.arr[key] = value }

func (this *MyHashMap) Get(key int) int { return this.arr[key] }

func (this *MyHashMap) Remove(key int) { this.arr[key] = -1 }

func main() {

}
