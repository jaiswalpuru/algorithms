package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	totalPeople = 23
	INTERVAL    = 1_00_000
)

func birthdaysList(n int) []int {
	list := make([]int, n)
	for i := 0; i < n; i++ {
		list[i] = rand.Intn(365)
	}
	return list
}

func returnUniqueList(list []int) []int {
	mp := make(map[int]bool)
	uniqueList := []int{}
	for i := 0; i < len(list); i++ {
		if _, ok := mp[list[i]]; !ok {
			mp[list[i]] = true
			uniqueList = append(uniqueList, list[i])
		}
	}
	return uniqueList
}

func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	var success = 0

	for i := 0; i < INTERVAL; i++ {
		birthdays := birthdaysList(totalPeople)
		uniqList := returnUniqueList(birthdays)
		if len(birthdays) != len(uniqList) {
			success++
		}
	}

	probability := float64(success) / float64(INTERVAL)
	fmt.Printf("Probability atleast 2 people in a group of %d, sharing birthdays is %.2f%%\n", totalPeople, 100.0*probability)
}
