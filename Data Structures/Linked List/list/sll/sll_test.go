package sll

import (
	"testing"
)

func TestSinglyLL(t *testing.T) {
	sll := New()

	sll.Insert(&Node{data: 1})
	sll.Insert(&Node{data: 2})
	sll.Insert(&Node{data: 3})
	sll.Insert(&Node{data: 4})
	sll.Insert(&Node{data: 5})

	res := sll.Search(3)
	if res == nil {
		t.Error("Insertion Error")
	}

	sll.Delete(res)
	if sll.Search(3) != nil {
		t.Error("Deletion error")
	}

}
