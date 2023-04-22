package main

type NestedInteger interface{}

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */

func depth_sum(l []*NestedInteger) int {
	return _depth_sum(l, 1)
}

func _depth_sum(l []*NestedInteger, depth int) int {
	sum := 0
	for i := 0; i < len(l); i++ {
		if l[i].IsInteger() {
			sum += (depth * l[i].GetInteger())
		} else {
			sum += _depth_sum(l[i].GetList(), depth+1)
		}
	}
	return sum
}

func depth_sum_bfs(l []*NestedInteger) int {
	sum := 0
	depth := 1
	q := []*NestedInteger{}
	q = append(q, l...)

	for len(q) > 0 {
		n := len(q)
		for i := 0; i < n; i++ {
			curr := q[0]
			q = q[1:]
			if curr.IsInteger() {
				sum += (depth * curr.GetInteger())
			} else {
				q = append(q, curr.GetList()...)
			}
		}
		depth++
	}
	return sum
}
