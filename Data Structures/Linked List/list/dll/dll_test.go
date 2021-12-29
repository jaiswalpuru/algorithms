package dll

import (
	"fmt"
	"testing"
)

func TestDoublyLL(t *testing.T) {
	dll := New()

	dll.Insert(&Node{data: 1})
	dll.Insert(&Node{data: 2})
	dll.Insert(&Node{data: 3})
	dll.Insert(&Node{data: 4})
	// fmt.Println(dll.Search(3))

	res := dll.Search(3)
	if res == nil {
		t.Error("Insertion error")
	}

	dll.Delete(res)

	itr := dll.head
	for itr != nil {
		fmt.Print(itr.data, " ")
		itr = itr.next
	}
	fmt.Println()
}
