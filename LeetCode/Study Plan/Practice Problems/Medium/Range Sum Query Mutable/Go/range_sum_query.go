package main

//-------------------------------Brute force approach-----------------------------------
type NumArray struct{ arr []int }

func Constructor(nums []int) NumArray { NumArray{arr: nums} }

func (this *NumArray) Update(index int, val int) { this.arr[index] = val }

func (this *NumArray) SumRange(left int, right int) int {
	sum := 0
	for i := left; i <= right; i++ {
		sum += this.arr[i]
	}
	return sum
}

//-------------------------------Brute force approach-----------------------------------

//-------------------------------Using segment tree approach-----------------------------------
type NumArray struct {
	seg_tree []int
	n        int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	tree := make([]int, n*2)
	//build the segment tree
	for i, j := n, 0; i < n*2; i++ {
		tree[i] = nums[j]
		j++
	}
	for i := n - 1; i > 0; i-- {
		tree[i] = tree[i*2] + tree[i*2+1]
	}
	return NumArray{seg_tree: tree, n: n}
}

func (this *NumArray) Update(index int, val int) {
	index += this.n
	this.seg_tree[index] = val
	for index > 0 {
		left := index
		right := index
		if index%2 == 0 {
			right = index + 1
		} else {
			left = index - 1
		}
		this.seg_tree[index/2] = this.seg_tree[left] + this.seg_tree[right]
		index /= 2
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	left += this.n
	right += this.n
	sum := 0
	for left <= right {
		if left%2 == 1 {
			sum += this.seg_tree[left]
			left++
		}
		if right%2 == 0 {
			sum += this.seg_tree[right]
			right--
		}
		left /= 2
		right /= 2
	}
	return sum
}

//-------------------------------Using segment tree approach-----------------------------------

func main() {

}
