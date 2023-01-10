package main

import "fmt"

type UnionFind struct{ root []int }

func Constructor(node int) UnionFind {
	var u UnionFind
	u.root = make([]int, node)
	for i := 0; i < node; i++ {
		u.root[i] = i
	}
	return u
}

func (this *UnionFind) union(x, y int) {
	root_x := this.find(x)
	root_y := this.find(y)
	if root_x != root_y {
		this.root[root_y] = root_x
	}
}

func (this *UnionFind) find(x int) int {
	for x != this.root[x] {
		x = this.root[x]
	}
	return x
}

func (this *UnionFind) is_connected(x, y int) bool { return this.find(x) == this.find(y) }

func main() {
	u := Constructor(10)
	u.union(1, 2)
	u.union(2, 5)
	u.union(5, 6)
	u.union(6, 7)
	u.union(3, 8)
	u.union(8, 9)
	fmt.Println(u.is_connected(1, 5))
	fmt.Println(u.is_connected(5, 7))
	fmt.Println(u.is_connected(4, 9))
	u.union(9, 4)
	fmt.Println(u.is_connected(4, 9))

}
