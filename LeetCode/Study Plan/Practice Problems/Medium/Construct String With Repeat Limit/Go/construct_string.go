package main

import (
	"container/heap"
	"fmt"
)

type Pair struct {
	char byte
	cnt  int
}

type MaxHeap []Pair

func (p MaxHeap) Len() int              { return len(p) }
func (p MaxHeap) Less(i, j int) bool    { return p[i].char > p[j].char }
func (p MaxHeap) Swap(i, j int)         { p[i], p[j] = p[j], p[i] }
func (p *MaxHeap) Push(val interface{}) { *p = append(*p, val.(Pair)) }
func (p *MaxHeap) Pop() interface{} {
	val := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return val
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func construct_string_with_repeat_limit(str string, repeat_limit int) string {
	count := make(map[byte]int)
	n := len(str)

	for i := 0; i < n; i++ {
		count[str[i]]++
	}

	mh := &MaxHeap{}
	for k, v := range count {
		heap.Push(mh, Pair{k, v})
	}

	s := ""
	for mh.Len() > 0 {
		curr := heap.Pop(mh).(Pair)
		c1 := curr.char
		cnt_c1 := curr.cnt

		t := min(repeat_limit, cnt_c1)
		for i := 0; i < t; i++ {
			s += string(c1)
		}

		if mh.Len() > 0 && cnt_c1-t > 0 {
			curr2 := heap.Pop(mh).(Pair)
			c2 := curr2.char
			cnt_c2 := curr2.cnt
			s += string(c2)
			if cnt_c1-t > 0 {
				heap.Push(mh, Pair{c1, cnt_c1 - t})
			}
			if cnt_c2-1 > 0 {
				heap.Push(mh, Pair{c2, cnt_c2 - 1})
			}
		}
	}
	return s
}

func main() {
	fmt.Println(construct_string_with_repeat_limit("cczazcc", 3))
}
