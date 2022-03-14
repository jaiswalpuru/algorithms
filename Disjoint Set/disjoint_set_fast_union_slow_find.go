package main

import "fmt"

type Element struct {
	parent *Element
	Data   interface{}
}

//Make set create a set of one element
func MakeSet(Data interface{}) *Element {
	s := &Element{}
	s.parent = s
	s.Data = Data
	return s
}

func Find(e *Element) *Element {
	for e.parent != e {
		e = e.parent
	}
	return e
}

//recursive
func FindRecursoce(e *Element) *Element {
	if e.parent == e {
		return e
	} else {
		return Find(e.parent)
	}
}

//Union establishes the union of two nodes
func Union(e1, e2 *Element) {
	e1.parent = e2
}

//test codes

func main() {
	aSet := MakeSet("a")
	bSet := MakeSet("b")
	oneSet := MakeSet(1)
	twoSet := MakeSet(2)

	Union(aSet, bSet)
	Union(oneSet, twoSet)

	result := Find(aSet)
	fmt.Println(result.Data)

	result = Find(oneSet)
	fmt.Println(result.Data)

	result = Find(twoSet)
	fmt.Println(result.Data)
}
