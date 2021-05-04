// Given an array of time intervals (start, end) for classroom lectures (possibly overlapping),
// find the minimum number of rooms required.

// For example, given [(30, 75), (0, 50), (60, 150)], you should return 2.

package main

import (
	"fmt"
	"sort"
)

func main() {

	//p := V{{30, 75}, {0, 50}, {60, 150}}
	startTime := []int{30, 0, 60}
	endTime := []int{75, 50, 150}
	sort.Ints(startTime)
	sort.Ints(endTime)
	rooms := 1
	onGoingMeetings := 1
	i, j := 1, 0

	for i < len(startTime) && j < len(startTime) {
		if startTime[i] < endTime[j] {
			onGoingMeetings++
			if onGoingMeetings > rooms {
				rooms = onGoingMeetings
			}
			i++
		} else {
			onGoingMeetings--
			j++
		}
	}

	fmt.Println(rooms)
}
