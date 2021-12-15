package main

//N-ary
const N = 3

type GenericTreeNode struct {
	data     int
	depth    int
	parent   *GenericTreeNode
	children []*GenericTreeNode
}

func (g *GenericTreeNode) addChildren() {
	for i := 0; i < N; i++ {
		newChild := &GenericTreeNode{
			parent: g,
			depth:  g.depth + 1,
		}
		g.children = append(g.children, newChild)
	}
}

func New() *GenericTreeNode {
	var root GenericTreeNode
	root.addChildren()

	for _, child := range root.children {
		child.addChildren()
	}
	return &root
}
