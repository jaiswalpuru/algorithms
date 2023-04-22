package main

func depthSumInverse(l []*NestedInteger) int {
	d := 1
	get_max_depth(l, &d, 1)
	type Pair struct {
		val   []*NestedInteger
		depth int
	}

	q := []Pair{{l, 1}}
	sum := 0
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		list := curr.val
		depth := curr.depth
		for i := 0; i < len(list); i++ {
			if list[i].IsInteger() {
				sum += (d - depth + 1) * list[i].GetInteger()
			} else {
				q = append(q, Pair{list[i].GetList(), depth + 1})
			}
		}
	}

	return sum
}

func get_max_depth(n []*NestedInteger, d *int, depth int) {
	for i := 0; i < len(n); i++ {
		if !n[i].IsInteger() {
			*d = max(*d, depth+1)
			get_max_depth(n[i].GetList(), d, depth+1)
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
