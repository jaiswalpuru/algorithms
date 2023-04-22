package main

import (
	"fmt"
	"sort"
	"strconv"
)

type Pair struct {
	val byte
	ind int
}

type I []Pair

func (p I) Len() int           { return len(p) }
func (p I) Less(i, j int) bool { return p[i].val > p[j].val }
func (p I) Swap(i, j int)      { p[i].val, p[j].val = p[j].val, p[i].val }

func largest_number_after_digit_swaps_by_parity(num int) int {
	str := []byte(strconv.Itoa(num))
	n := len(str)

	odd, even := I{}, I{}
	for i := 0; i < n; i++ {
		if int(str[i]-'0')%2 == 0 {
			even = append(even, Pair{str[i], i})
		} else {
			odd = append(odd, Pair{str[i], i})
		}
	}

	sort.Sort(even)
	sort.Sort(odd)

	res := make([]byte, n)
	for i := 0; i < len(even); i++ {
		res[even[i].ind] = even[i].val
	}

	for i := 0; i < len(odd); i++ {
		res[odd[i].ind] = odd[i].val
	}

	val, _ := strconv.Atoi(string(res))
	return val
}

func main() {
	fmt.Println(largest_number_after_digit_swaps_by_parity(1234))
}
