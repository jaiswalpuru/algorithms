package main

import "fmt"
import "sort"

type Pair struct { rock, cap int }

type P []Pair

func abs(a int) int {
	if a< 0 {
		return -1 * a
	}
	return a
}

func (p P) Len() int {return len(p)}
func (p P) Swap(i,j int) {p[i],p[j] = p[j],p[i]}
func(p P) Less(i,j int)bool {return (p[i].cap-p[i].rock) < (p[j].cap-p[j].rock)}

func max_bag_with_full_capacity_of_rocks(capacity, rocks []int, additional_rocks int) int {
	n := len(capacity)
	arr := make(P, n)
	for i:=0;i<n;i++{
		arr[i].rock = rocks[i]
		arr[i].cap = capacity[i]
	}

	cnt := 0
	sort.Sort(arr)
	for i :=0; i<n;i++{
		reqd := arr[i].cap-arr[i].rock
		if reqd > 0 && additional_rocks > 0{
			if additional_rocks >= reqd {
				additional_rocks -= reqd
				cnt++
			}else {
				additional_rocks = reqd - additional_rocks
			}
		}
		if arr[i].cap-arr[i].rock == 0 {
			cnt++
		}
	}

	return cnt
}

func main () {
	fmt.Println(max_bag_with_full_capacity_of_rocks([]int{10,2,2}, []int{2,2,0}, 100))
}