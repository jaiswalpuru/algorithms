package main

type Element struct {
	parent *Element
	rank   int
	Data   interface{}
}

func MakeSet(Data interface{}) *Element {
	s := &Element{}
	s.parent = s
	s.rank = 1
	s.Data = Data
	return s
}

func Find(e *Element) *Element {
	for e != e.parent {
		e = e.parent
	}
	return e
}

//Unionbyheight
//In this method we store the height of the tree
func Union(e1, e2 *Element) {
	//Ensure the two elements are not already the part of the same union
	e1SetName := Find(e1)
	e2SetName := Find(e2)
	if e1SetName == e2SetName {
		return
	}

	//Create a union by making the shorter tree point to the root of the larger tree
	switch {
	case e1SetName.rank < e2SetName.rank:
		e1SetName.parent = e2SetName
	case e1SetName.rank > e2SetName.rank:
		e2SetName.parent = e1SetName
	default:
		e2SetName.parent = e1SetName
		e1SetName.rank++
	}
}

func main() {

}
