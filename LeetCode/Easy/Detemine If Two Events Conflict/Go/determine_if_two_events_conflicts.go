package main

import "fmt"

func determine_if_two_events_conflicts(e1, e2 []string) bool {
	if e1[0] >= e2[0] && e1[0] <= e2[1] {
		return true
	}
	if e2[0] >= e1[0] && e2[0] <= e1[1] {
		return true
	}
	return false
}

func main() {
	fmt.Println(determine_if_two_events_conflicts([]string{"01:15", "02:00"}, []string{"02:00", "03:00"}))
}
