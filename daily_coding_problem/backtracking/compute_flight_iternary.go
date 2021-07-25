/*
Given an unoredered list of flights taken by someone, each represented as (origin, destination) pairs, and
a starting airport, compute a possible iternary. If no such iternary exists, return null. All flights must be
used in the iternary.
*/

package main

import "fmt"

func _intinary(pair [][]string, itirnary []string) []string {
	if len(pair) == 0 {
		return itirnary
	}
	last_stop := itirnary[len(itirnary)-1]
	for i := range pair {
		temp := [][]string{}
		//copying all the flights except the current one
		temp = append(temp, pair[:i]...)
		temp = append(temp, pair[i+1:]...)
		itirnary = append(itirnary, pair[i][1])
		if pair[i][0] == last_stop {
			//fmt.Println(i, pair[i][0], temp)
			return _intinary(temp, itirnary)
		}
		itirnary = itirnary[:len(itirnary)-1]
	}
	return nil
}

func get_itirnary(pair [][]string, start string) []string {
	current_iternary := make([]string, 0)
	current_iternary = append(current_iternary, start)
	return _intinary(pair, current_iternary)
}

func main() {
	pair := [][]string{
		{"SFO", "HKO"},
		{"YYZ", "SFO"},
		{"YUL", "YYZ"},
		{"HKO", "ORD"},
	}
	fmt.Println(get_itirnary(pair, "YUL"))
	pair_one := [][]string{
		{"SFO", "COM"},
		{"COM", "YYZ"},
	}
	fmt.Println(get_itirnary(pair_one, "COM"))
}
