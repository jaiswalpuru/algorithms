package main

import (
	"errors"
	"fmt"
)

const (
	inf = int(^uint(0) >> 1)
)

var ErrNegativeCycle = errors.New("bellman : graph contains negative cycle")
var ErrNoPath = errors.New("bellman : no path vertex")

type Edges []*Edge

type Edge struct {
	From, To string
	Distance int
}

type Vertices map[string]*Vertex

type Vertex struct {
	Distance int
	From     string // predecsessor
}

//ShortestPath finds the shortest path form the source to destination after calling Search
func (path Vertices) ShortestPath(src, dst string) ([]*Vertex, error) {
	if len(path) == 0 {
		//called shortest path before calling the Search
		return nil, ErrNoPath
	}
	var all []*Vertex
	pred := dst
	for pred != src {
		if v := path[pred]; v != nil {
			pred = v.From
			all = append(all, v)
		} else {
			return nil, ErrNoPath
		}
	}
	return all, nil
}

//Add Edge for search
func (rt *Edges) AddEdge(from, to string, distance int) {
	*rt = append(*rt, &Edge{From: from, To: to, Distance: distance})
}

//Search for single source shortest path
func (rt Edges) Search(start string) (Vertices, error) {
	//resulting table contains vertex name to vertex struct mapping
	//use v.From predecessor to trace the path back
	distances := make(Vertices)
	for _, p := range rt {
		distances[p.From] = &Vertex{inf, " "}
		distances[p.To] = &Vertex{inf, " "}
	}

	distances[start].Distance = 0
	//As many iterations as there are nodes
	for i := 0; i < len(distances); i++ {
		var changed bool
		//iterate pairs
		for _, pair := range rt {
			n := distances[pair.From]
			if n.Distance != inf {
				if distances[pair.To].Distance > n.Distance+pair.Distance {
					distances[pair.To].Distance = n.Distance + pair.Distance
					distances[pair.To].From = pair.From
					changed = true
				}
			}
		}
		if !changed {
			//no more changes made
			break
		}
	}

	for _, pair := range rt {
		if distances[pair.From].Distance != inf && distances[pair.To].Distance != inf && distances[pair.From].Distance+pair.Distance < distances[pair.To].Distance {
			return nil, ErrNegativeCycle
		}
	}

	return distances, nil
}

type testcase struct {
	src, dest string
	path      string
}

var tests = []*testcase{
	{"Hyderabad", "Chennai", "Hyderabad 106,Bangalore 95,"},
	{"Bangalore", "Mumbai", "Bangalore 95,Chennai 65,Canada 73,Pune 62,Amaravathi 92,"},
	{"Delhi", "Chennai", "Delhi 121,"},
}

func main() {
	fmt.Println("Bellman Ford")
	for _, test := range tests {
		rt := Edges{
			{"Hyderabad", "Vijayawada", 117},
			{"Vijayawada", "Vizag", 65},
			{"Vizag", "Pune", 70},
			{"Pune", "Amaravathi", 62},
			{"Amaravathi", "Mumbai", 92},
			{"Hyderabad", "Bangalore", 106},
			{"Hyderabad", "Bangalore", 113},
			{"Bangalore", "Chennai", 95},
			{"Vijayawada", "Guntur", 24},
			{"Guntur", "Chennai", 88},
			{"Vizag", "Chennai", 94},
			{"Chennai", "Canada", 65},
			{"Canada", "Pune", 73},
			{"Chennai", "Delhi", 121},
			{"Canada", "Delhi", 103},
		}
		for _, r := range rt {
			// Add reversed edges.
			rt = append(rt, &Edge{From: r.To, To: r.From, Distance: r.Distance})
		}
		//fmt.Println(rt)
		paths, err := rt.Search(test.src)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(paths)
		all, err := paths.ShortestPath(test.src, test.dest)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(all)
		var prev int
		var s string
		for i := len(all) - 1; i >= 0; i-- {
			s += fmt.Sprintf("%s %d,", all[i].From, all[i].Distance-prev)
			prev = all[i].Distance
		}

		if test.path != s {
			fmt.Printf("shortest path does not match, got %q want %q", s, test.path)
		}
		fmt.Println(s)
	}
}
