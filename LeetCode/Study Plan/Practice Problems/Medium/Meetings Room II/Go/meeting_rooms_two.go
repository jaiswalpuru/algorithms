package main

import ("fmt"; "sort"; "container/heap")

type P []int
func (p P) Len()int {return len(p)}
func (p P) Less(i,j int) bool {return p[i]<p[j]}
func (p P) Swap(i,j int) {p[i],p[j] = p[j],p[i]}
func (p *P) Push(val interface{}) {*p = append(*p, val.(int))}
func (p *P) Pop() interface{} {
	v := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return v
}

type I [][]int
func (a I)Len() int {return len(a)}
func  (a I) Less(i,j int) bool {
	return a[i][0] < a[j][0]
}
func (a I) Swap(i,j int) {a[i],a[j] = a[j],a[i]}

func meeting_rooms_two(arr [][]int) int {
	var a I
	a = arr
	sort.Sort(a)

	mh := &P{}

	rooms := 1
	heap.Push(mh, a[0][1])
	for i:=1; i<a.Len(); i++{
		val := heap.Pop(mh).(int)
		if val > a[i][0] {
			rooms++
			heap.Push(mh, val)
		}
		heap.Push(mh, a[i][1])
	}

	return rooms
}

func main() {
	fmt.Println(meeting_rooms_two([][]int{{0,30}, {5,10}, {15,20}}))
}