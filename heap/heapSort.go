package main

//HeapSort - a sorting technique which uses heap and Time Complexity : O(nlogn)
func HeapSort(data []Item) []Item {
	hp := Heapify(data)
	size := len(hp.data)
	for i := size - 1; i > 0; i-- {
		//Move root to end
		swap(hp, 0, i)
		hp.size--
		hp.percolateDown(0)
	}
	hp.size = size
	return hp.data
}
