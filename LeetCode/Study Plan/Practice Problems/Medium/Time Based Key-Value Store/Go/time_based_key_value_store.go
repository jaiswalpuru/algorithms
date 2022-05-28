package main

type Pair struct {
	val      string
	time_stamp int
}

type TimeMap struct{ tree_map map[string][]Pair }

func Constructor() TimeMap {
    return TimeMap{tree_map : make(map[string][]Pair)}
}

func (this *TimeMap) Set(key, val string, time_stamp int) {
	this.tree_map[key] = append(this.tree_map[key], Pair{val, time_stamp})
}

func b_search(arr []Pair, val int) int {
	l, h := 0, len(arr)-1
	res := -1
	for l <= h {
		mid := (l + h) / 2
		if arr[mid].time_stamp <= val {
			l = mid + 1
			res = mid
		} else {
			h = mid - 1
		}
	}
	return res
}

func (this *TimeMap) Get(key string, time_stamp int) string {
	arr := this.tree_map[key]
	if arr == nil {
		return ""
	}

	ind := b_search(arr, time_stamp)
	if ind == -1 {
		return ""
	}
	return arr[ind].val
}
