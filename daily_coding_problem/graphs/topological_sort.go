package main

import "fmt"

func topolocial_sort(courses map[int][]int) []int {
	topo_order := []int{}
	in_degree := make(map[int]int)

	//Set all the indegree values to zero
	for u := range courses {
		in_degree[u] = 0
	}

	//For each vertex u
	for _, adjacenct := range courses {
		//For each vertex adjacent to u
		for _, v := range adjacenct {
			in_degree[v]++
		}
	}

	//make a list of all vertices which have indegree as zero
	next := []int{}
	for u, v := range in_degree {
		if v != 0 {
			continue
		}
		next = append(next, u)
	}

	for len(next) > 0 {
		//process the vertex for start which have indegree as zero
		u := next[0]
		next = next[1:]

		//add to the topo_order
		topo_order = append(topo_order, u)
		for _, v := range courses[u] {
			//decrement the indegree for each vertex which have pre_req as u
			in_degree[v]--
			if in_degree[v] == 0 {
				next = append(next, v)
			}
		}
	}

	return topo_order
}

func main() {
	// pre_req := map[string][]string{
	// 	"CSC300": []string{"CSC100", "CSC200"},
	// 	"CSC200": []string{"CSC100"},
	// 	"CSC100": []string{},
	// }

	courses := map[int][]int{
		1:  {4},
		2:  {3},
		3:  {4, 5},
		4:  {6},
		5:  {6},
		6:  {7, 11},
		7:  {8},
		8:  {14},
		9:  {10},
		10: {11},
		11: {12},
		12: {13},
		13: {},
	}

	fmt.Println(topolocial_sort(courses))

}
