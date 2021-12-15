package binary_heaps

import (
	"testing"
)

type Int int

func (a Int) Less(b Item) bool {
	val, ok := b.(Int)
	return ok && a <= val
}

func TestHeap(t *testing.T) {
	arr := []Item{Int(27), Int(17), Int(3), Int(16), Int(13), Int(10), Int(1), Int(5), Int(7), Int(12), Int(4), Int(8), Int(9), Int(0)}
	result := Heapify(arr)
	for i := 0; i < result.size; i++ {
		left := left_child(i)
		right := right_child(i)

		if left >= result.size || right >= result.size {
			break
		}

		if !result.data[i].Less(result.data[left]) || !result.data[i].Less(result.data[right]) {
			t.Error("Error in heapify")
		}
	}

}
