package main

//-----------------brute force---------------------------
type Pair struct{ s, e int }

type MyCalendar struct{ events []Pair }

func Constructor() MyCalendar { return MyCalendar{events: make([]Pair, 0)} }

func (this *MyCalendar) Book(s int, e int) bool {
	for i := 0; i < len(this.events); i++ {
		if this.events[i].s < e && this.events[i].e > s {
			return false
		}
	}
	this.events = append(this.events, Pair{s, e})
	return true
}

//-----------------brute force---------------------------

//---------------------using binary tree-----------------------
type Node struct {
	start,
	end int
	left  *Node
	right *Node
}

func New(s, e int) *Node { return &Node{start: s, end: e, left: nil, right: nil} }

func insert(root, node *Node) bool {
	if root.end <= node.start {
		if root.right == nil {
			root.right = node
			return true
		}
		return insert(root.right, node)
	} else if root.start >= node.end {
		if root.left == nil {
			root.left = node
			return true
		}
		return insert(root.left, node)
	} else {
		return false
	}
}

type MyCalendar struct{ node *Node }

func Constructor() MyCalendar { return MyCalendar{node: nil} }

func (this *MyCalendar) Book(s, e int) bool {
	if this.node == nil {
		this.node = New(s, e)
		return true
	}
	return insert(this.node, New(s, e))
}

//---------------------using binary tree-----------------------
