//below is the brute force solution

type SummaryRanges struct {
	intervals map[int]bool
}

func Constructor() SummaryRanges {
	return SummaryRanges{intervals: make(map[int]bool)}
}

func (this *SummaryRanges) AddNum(value int) {
	this.intervals[value] = true
}

func (this *SummaryRanges) GetIntervals() [][]int {
	l, r := -1, -1
	res := [][]int{}
	for i := 0; i < int(1e4); i++ {
		if this.intervals[i] {
			if l == -1 {
				l, r = i, i
			} else {
				r = i
			}
		} else {
			if l != -1 && r != -1 {
				res = append(res, []int{l, r})
			}
			l, r = -1, -1
		}
	}
	return res
}

/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(value);
 * param_2 := obj.GetIntervals();
 */
