package main

import "fmt"

//-------------------------------Brute force approach-----------------------------------
type NumArray struct { arr []int }


func Constructor(nums []int) NumArray {
    n := len(nums)
    var n_a NumArray
    n_a.arr = make([]int, n)
    for i:=0; i<n; i++{
    	n_a.arr[i] = nums[i]
    }
    return n_a
}


func (this *NumArray) Update(index int, val int)  { this.arr[index] = val }


func (this *NumArray) SumRange(left int, right int) int {
    sum := 0
    for i:=left; i<=right; i++{
    	sum += this.arr[i]
    }
    return sum
}
//-------------------------------Brute force approach-----------------------------------


//-------------------------------Using segment tree approach-----------------------------------
type NumArray struct {
	tree []int
	n int
}

func build_tree(tree *[]int, nums []int) {
    n := len(nums)
	for i,j := n, 0; i<n*2; i++{
		(*tree)[i] = nums[j]
		j++
	}

	for i:=n-1; i>0;i-- {
		(*tree)[i] = (*tree)[i*2] + (*tree)[i*2+1]
	}
}

func Constructor(nums []int) NumArray {
	var n_a NumArray
	n := len(nums)
	if n > 0 {
		tree := make([]int, n*2)
		build_tree(&tree, nums)
        n_a.tree = tree
	    n_a.n = len(nums)
	}	
	return n_a
}


func (this *NumArray) Update(index int, val int)  { 
	//when we are updating the index we need to rebuild the segment tree and because of other tree nodes
	//which contains sum of the modified element.
	/*
		1. First we update the position at leaf node.
		2. And from there again we use bottom up approach to update the indices
	*/
	index += this.n
	this.tree[index] = val

	for index > 0 {
		left := index
		right := index
		if index % 2 == 0 {
			right = index + 1
		} else {
			left = index - 1
		}
		this.tree[index/2] = this.tree[left] + this.tree[right]
		index /= 2
	}
}


func (this *NumArray) SumRange(left int, right int) int {
    l := this.n + left
    r := this.n + right
    sum := 0
    for l<=r {
    	if l%2 == 1{
    		sum += this.tree[l]
    		l++
    	}
    	if r%2 == 0 {
    		sum += this.tree[r]
    		r--
    	}
    	l/=2
    	r/=2
    }
    return sum
}
//-------------------------------Using segment tree approach-----------------------------------


func main() {
	
}