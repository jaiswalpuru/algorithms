/*
Author : Puru Jaiswal
References : CLRS
*/

package stack

import "fmt"

type Item interface{}

type Stack struct {
	data []Item
	top  int
}

func New() *Stack { return &Stack{top: -1} }

func (s *Stack) Is_Empty() bool { return s.top == -1 }

//Push
func (s *Stack) Push(val Item) {
	if s.top == -1 {
		s.data = make([]Item, 1)
		s.top++
		s.data[s.top] = val
	} else {
		s.data = append(s.data, val)
		s.top++
	}
}

//Pop
func (s *Stack) Pop() Item {
	if s.top == -1 {
		fmt.Println("Stack empty")
		return nil
	}
	data := s.data[s.top]
	if s.top == 0 {
		s.top = -1
		s.data = []Item{}
	} else {
		s.data = s.data[:s.top-1]
		s.top--
	}
	return data
}
