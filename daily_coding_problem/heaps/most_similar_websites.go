package main

import (
	"container/heap"
	"fmt"
)

type Pair struct {
	website string
	user    int
}

type P struct {
	site_one string
	site_two string
}

type Triplet struct {
	score float64
	pair  P
}

type MinHeap []Triplet

func (mh MinHeap) Len() int { return len(mh) }
func (mh MinHeap) Swap(i, j int) {
	temp_one, temp_two := mh[i], mh[j]
	mh[j].pair = temp_one.pair
	mh[j].score = temp_one.score
	mh[i].pair = temp_two.pair
	mh[i].score = temp_two.score
}
func (mh MinHeap) Less(i, j int) bool { return mh[i].score > mh[j].score }

func (mh *MinHeap) Push(val interface{}) { *mh = append(*mh, val.(Triplet)) }

func (mh *MinHeap) Pop() interface{} {
	old := *mh
	n := len(old)
	val := old[n-1]
	*mh = old[:n-1]
	return val
}

func main() {
	k := 1
	users := []Pair{
		{"google.com", 1}, {"google.com", 3}, {"google.com", 5},
		{"pets.com", 1}, {"pets.com", 2}, {"pets.com", 6},
		{"yahoo.com", 2}, {"yahoo.com", 3}, {"yahoo.com", 4}, {"yahoo.com", 5},
		{"wikipedia.org", 4}, {"wikipedia.org", 5}, {"wikipedia.org", 6}, {"wikipedia.org", 7},
		{"bing.com", 1}, {"bing.com", 3}, {"bing.com", 5}, {"bing.com", 6},
	}

	mp := make(map[string][]int)

	for i := 0; i < len(users); i++ {
		mp[users[i].website] = append(mp[users[i].website], users[i].user)
	}

	site_pair := make(MinHeap, 0)
	sites := []string{}
	for k := range mp {
		sites = append(sites, k)
	}

	heap.Init(&site_pair)

	for i := 0; i < k; i++ {
		heap.Push(&site_pair, Triplet{score: 0, pair: P{site_one: "", site_two: ""}})
	}

	for i := 0; i < len(sites)-1; i++ {
		for j := i + 1; j < len(sites); j++ {
			score := compute_score(sites[i], sites[j], mp)
			heap.Push(&site_pair, Triplet{score: score, pair: P{site_one: sites[i], site_two: sites[j]}})
			//heap.Pop(&site_pair)
		}
	}

	list := []P{}
	for i := 0; i < len(site_pair); i++ {
		list = append(list, site_pair[i].pair)
	}
	fmt.Println(list[0])
}

func compute_score(site_one, site_two string, visitors map[string][]int) float64 {
	x, y := get_common_elements(visitors[site_one], visitors[site_two])
	return float64(x) / float64(y)
}

func get_common_elements(a, b []int) (int, int) {
	if len(b) > len(a) {
		a, b = b, a
	}
	temp := make(map[int]bool)
	for i := 0; i < len(a); i++ {
		temp[a[i]] = true
	}

	common := 0
	for i := 0; i < len(b); i++ {
		if _, ok := temp[b[i]]; ok {
			common++
		}
	}
	return common, len(a) + len(b) - common
}
