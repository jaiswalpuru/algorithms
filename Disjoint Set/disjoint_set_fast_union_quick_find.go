package main

type Element struct {
	parent *Element
	size   int
	Data   interface{}
}

func MakeSet(Data interface{}) *Element {
	s := &Element{}
	s.parent = s
	s.size = 1
	s.Data = Data
	return s
}

func Find(e *Element) *Element {
	for e != e.parent {
		e = e.parent
	}
	return e
}

func Union(e1, e2 *Element) {
	//Ensure the two elelments aren't already part of the same union
	e1SetName := Find(e1)
	e2SetName := Find(e2)
	if e1SetName == e2SetName {
		return
	}
	//Create a union by making the shorter tree point to the root of the larger tree
	if e1SetName.size < e2SetName.size {
		e1SetName.parent = e2SetName
		e2SetName.size += e1SetName.size
	} else {
		e2SetName.parent = e1SetName
		e1SetName.size += e2SetName.size
	}
}

//FindWithPathCompression
func FindWithPathCompression(e *Element) *Element {
	for e != e.parent {
		e.parent = e.parent.parent
		e = e.parent
	}
	return e
}

func FindwithPathCompressionRecursive(e *Element) *Element {
	if e.parent == e {
		return e
	}
	
	e.parent = return FindwithPathCompressionRecursive(e.parent)
	return e.parent
}

//there is no change in the find operation implementation
