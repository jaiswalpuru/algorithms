package main

import (
	"errors"
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

func main() {

}
