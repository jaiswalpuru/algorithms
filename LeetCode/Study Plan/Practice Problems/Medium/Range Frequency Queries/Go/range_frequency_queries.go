package main

//-------------------------------using segment tree--------------------------------------
type RangeFreqQuery struct {
	tree []map[int]int
	n    int
}

func build_tree(arr []int, index, low, high int, tree *[]map[int]int) {
	if low == high {
		(*tree)[index][arr[low]] = 1
		return
	}

	mid := (low + high) / 2
	left := 2 * index
	right := 2*index + 1
	build_tree(arr, left, low, mid, tree)
	build_tree(arr, right, mid+1, high, tree)
	for k := range (*tree)[left] {
		(*tree)[index][k] += (*tree)[left][k]
	}
	for k := range (*tree)[right] {
		(*tree)[index][k] += (*tree)[right][k]
	}
}

func Constructor(arr []int) RangeFreqQuery {
	n := len(arr)
	tree := make([]map[int]int, 4*n)
	for i := 0; i < 4*n; i++ {
		tree[i] = make(map[int]int)
	}
	build_tree(arr, 1, 0, len(arr)-1, &tree)
	var rfq RangeFreqQuery
	rfq.tree = tree
	rfq.n = n
	return rfq
}

func (this *RangeFreqQuery) Query(left int, right int, value int) int {
	return _query(1, 0, this.n-1, left, right, value, this.tree)
}

func _query(ind, low, high, left, right int, value int, tree []map[int]int) int {
	if right < low || left > high {
		return 0
	}

	if low >= left && high <= right {
		return tree[ind][value]
	}

	mid := (low + high) / 2
	_left := _query(2*ind, low, mid, left, right, value, tree)
	_right := _query(2*ind+1, mid+1, high, left, right, value, tree)
	return _left + _right
}

//-------------------------------using segment tree--------------------------------------

//-------------------------------using hashmap + binary search--------------------------------------

type RangeFreqQuery struct {
	hash_map map[int][]int
}

func Constructor(arr []int) RangeFreqQuery {
	hash_map := make(map[int][]int)
	n := len(arr)
	for i := 0; i < n; i++ {
		hash_map[arr[i]] = append(hash_map[arr[i]], i)
	}

	var rfq RangeFreqQuery
	rfq.hash_map = hash_map
	return rfq
}

func (this *RangeFreqQuery) Query(left, right, val int) int {
	list := this.hash_map[val]
	if list == nil {
		return 0
	}

	l := sort.Search(len(list), func(i int) bool { return list[i] >= left })
	r := sort.Search(len(list), func(i int) bool { return list[i] > right })

	return r - l
}

//-------------------------------using hashmap + binary search--------------------------------------
