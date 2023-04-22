package main

type Node struct {
	Val int
	Left,
	Right,
	Next *Node
}

func New(val int) *Node { return &Node{Val: val, Left: nil, Right: nil, Next: nil} }

func populate_next_right_pointers(root *Node) *Node {
	head := root
	q := []*Node{head}
	for len(q) > 0 {
		n := len(q)
		var curr *Node
		for i := 0; i < n; i++ {
			curr = q[0]
			q = q[1:]
			curr.Next = q[0]
			if curr.Left != nil {
				q = append(q, curr.Left)
			}
			if curr.Right != nil {
				q = append(q, curr.Right)
			}
		}
	}
	return root
}

//-----------------------------using previously established next pointers-----------------------------
func populate_next_pointers(root *Node) *Node {
	if root == nil {
		return nil
	}

	left := root
	for left.Left != nil {
		head := left
		for head != nil {
			head.Left.Next = head.Right
			if head.Next != nil {
				head.Right.Next = head.Next.Left
			}
			head = head.Next
		}
		left = left.Left
	}
	return root
}

//-----------------------------------------------------------------------------------------------------

func main() {

}
