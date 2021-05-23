/*
	Given an unordered list of flights taken by someone, each represented as (origin, destination) pairs,
	and a starting airport, compute the person's itinerary.
	If no such itinerary exists, return null. If there are multiple possible itineraries,
	return the lexicographically smallest one. All flights must be used in the itinerary.

	For example, given the list of flights [('SFO', 'HKO'), ('YYZ', 'SFO'), ('YUL', 'YYZ'), ('HKO', 'ORD')]
	and starting airport 'YUL', you should return the list ['YUL', 'YYZ', 'SFO', 'HKO', 'ORD'].

	Given the list of flights [('SFO', 'COM'), ('COM', 'YYZ')] and starting airport 'COM', you should return null.

	Given the list of flights [('A', 'B'), ('A', 'C'), ('B', 'C'), ('C', 'A')]
	and starting airport 'A', you should return the list ['A', 'B', 'C', 'A', 'C'] even
	though ['A', 'C', 'A', 'B', 'C'] is also a valid itinerary. However, the first one is lexicographically smaller.
*/

/*
	In several other PL's a datastructure set can be used as a value to the key
	--> map[string]set()
	so overhead of sorting won't be involved.
	Can also be solved using topological sorting.
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	visited := make(map[string]bool)
	back_track := make(map[string][]string)

	type pair struct {
		src, dst string
	}

	itenary := []pair{
		{"SFO", "COM"},
		{"COM", "YYZ"},
	}

	for i := 0; i < len(itenary); i++ {
		if _, ok := back_track[itenary[i].src]; ok {
			back_track[itenary[i].src] = append(back_track[itenary[i].src], itenary[i].dst)
		} else {
			visited[itenary[i].src] = true
			back_track[itenary[i].src] = make([]string, 0)
			back_track[itenary[i].src] = append(back_track[itenary[i].src], itenary[i].dst)
		}
	}

	starting_point := "YUL"

	for k, v := range back_track {
		if len(v) > 1 {
			sort.Strings(back_track[k])
		}
	}

	iternaries := []string{}
	iternaries = append(iternaries, starting_point)
	for {
		if len(back_track[starting_point]) == 0 {
			break
		}
		temp := back_track[starting_point][0]
		iternaries = append(iternaries, back_track[starting_point][0])
		back_track[starting_point] = back_track[starting_point][1:]
		starting_point = temp
	}

	if len(iternaries) >= len(visited) {
		fmt.Println(iternaries)
	} else {
		fmt.Println("break in the iternary")
	}
}
