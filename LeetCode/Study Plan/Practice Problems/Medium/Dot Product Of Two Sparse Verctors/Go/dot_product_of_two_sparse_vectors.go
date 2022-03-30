package main

type SparseVector struct {
	hash_map map[int]int
}

func Constructor(nums []int) SparseVector {
	var sv SparseVector
	hash_map := make(map[int]int)
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			hash_map[i] = nums[i]
		}
	}
	sv.hash_map = hash_map
	return sv
}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {
	sum := 0
	for k, v := range this.hash_map {
		if vec.hash_map[k] != 0 {
			sum += v * vec.hash_map[k]
		}
	}
	return sum
}

func main() {

}
