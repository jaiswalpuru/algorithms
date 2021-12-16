package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	st := New()
	st.Push(12)
	if st.data[st.top] != 12 {
		t.Error("Error in insertion")
	}

	d := st.Pop()
	if d != 12 {
		t.Error("Wrong Insertion")
	}

	if st.Is_Empty() {
		t.Log("Stack is empty")
	}
}
