package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	scanf  = func(f string, args ...interface{}) { fmt.Fscanf(reader, f, args...) }
	printf = func(f string, args ...interface{}) { fmt.Fprintf(writer, f, args...) }
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type xSorter []string

func (a xSorter) Len() int           { return len(a) }
func (a xSorter) Less(i, j int) bool { return compareString(a[i], a[j]) }
func (a xSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func compareString(a, b string) bool {
	lenA := len(a)
	lenB := len(b)
	if lenA > lenB {
		return false
	} else if lenA < lenB {
		return true
	} else {
		if a > b {
			return true
		}
		return false
	}
}

func main() {
	defer writer.Flush()

	var (
		n, m     int
		na, cnty string
	)
	scanf("%d %d\n", &n, &m)

	nameCnty := make(map[string]string)
	name := make([]string, m)
	votes := make(map[string]int)
	cntyVote := make(map[string]int)

	maxVote := -1
	maxVCnty := -1

	for i := 0; i < n; i++ {
		scanf("%s %s\n", &na, &cnty)
		nameCnty[na] = cnty
	}

	for i := 0; i < m; i++ {
		scanf("%s\n", &name[i])
		votes[name[i]]++
		maxVote = max(maxVote, votes[name[i]])
		cntyVote[nameCnty[name[i]]]++
		maxVCnty = max(maxVCnty, cntyVote[nameCnty[name[i]]])
	}

	//names of chef with max vote
	maxVoteChef := []string{}
	maxVoteCnty := []string{}
	for k, v := range votes {
		if v == maxVote {
			maxVoteChef = append(maxVoteChef, k)
		}
	}

	for k, v := range cntyVote {
		if maxVCnty == v {
			maxVoteCnty = append(maxVoteCnty, k)
		}
	}

	sort.Strings(maxVoteCnty)
	sort.Strings(maxVoteChef)

	fmt.Println(maxVoteCnty[0])
	fmt.Println(maxVoteChef[0])
}
