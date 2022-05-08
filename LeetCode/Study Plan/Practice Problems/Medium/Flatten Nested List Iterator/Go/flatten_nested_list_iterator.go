package main

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */
func get_list_val(l []*NestedInteger) []int {
    list := []int{}
    for i:=0;i<len(l);i++{
        if l[i].IsInteger() {
            list = append(list, l[i].GetInteger())
        }else {
            t := get_list_val(l[i].GetList())
            list = append(list, t...)
        }
    }
    return list
}

type NestedIterator struct { 
    l []int
    i int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
    l := get_list_val(nestedList)
    return &NestedIterator{i:0, l:l}
}

func (this *NestedIterator) Next() int {
    val := this.l[this.i]
    this.i+=1
    return val
}

func (this *NestedIterator) HasNext() bool {
    return this.i<len(this.l)
}