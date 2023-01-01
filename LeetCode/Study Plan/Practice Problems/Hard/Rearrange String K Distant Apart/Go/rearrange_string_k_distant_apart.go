package main

import (
	"container/heap"
	"fmt"
)

type P struct {
	char byte
	cnt  int
}

type M []P

func (m M) Len() int              { return len(m) }
func (m M) Less(i, j int) bool    { return m[i].cnt > m[j].cnt }
func (m M) Swap(i, j int)         { m[i], m[j] = m[j], m[i] }
func (m *M) Push(val interface{}) { *m = append(*m, val.(P)) }
func (m *M) Pop() interface{} {
	val := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return val
}

func rearrange_string_k_distant_apart(s string, k int) string {
	hash_map := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		hash_map[s[i]]++
	}
	mh := &M{}
	for k, v := range hash_map {
		heap.Push(mh, P{k, v})
	}
	q := []P{}
	res := ""
	for mh.Len() > 0 {
		curr := heap.Pop(mh).(P)
		res += string(curr.char)
		curr.cnt -= 1
		q = append(q, curr)
		if len(q) < k {
			continue
		}
		front := q[0]
		q = q[1:]
		if front.cnt > 0 {
			heap.Push(mh, front)
		}
	}
	if len(res) != len(s) {
		return ""
	}
	return res
}

func main() {
	fmt.Println(rearrange_string_k_distant_apart("aabbcc", 3))
}
