package main

import "fmt"

func make_graph(recipes []string, ingridient [][]string) map[string][]string {
	n := len(recipes)
	graph := make(map[string][]string)
	for i := 0; i < n; i++ {
		for j := 0; j < len(ingridient[i]); j++ {
			if _, ok := graph[ingridient[i][j]]; !ok {
				graph[ingridient[i][j]] = make([]string, 0)
			}
			graph[ingridient[i][j]] = append(graph[ingridient[i][j]], recipes[i])
		}
	}
	return graph
}

func recipes(recipes []string, ingredients [][]string, supplies []string) []string {
	g := make_graph(recipes, ingredients)
	indegree := make(map[string]int)
	n := len(recipes)
	for i := 0; i < n; i++ {
		indegree[recipes[i]] = len(ingredients[i])
	}

	recipes_with_indegree_zero := []string{}
	for i := 0; i < len(supplies); i++ {
		recipes_with_indegree_zero = append(recipes_with_indegree_zero, supplies[i])
	}

	res := []string{}
	for len(recipes_with_indegree_zero) > 0 {
		curr := recipes_with_indegree_zero[0]
		recipes_with_indegree_zero = recipes_with_indegree_zero[1:]
		for i := 0; i < len(g[curr]); i++ {
			indegree[g[curr][i]]--
			if indegree[g[curr][i]] == 0 {
				res = append(res, g[curr][i])
				recipes_with_indegree_zero = append(recipes_with_indegree_zero, g[curr][i])
			}
		}
	}

	return res
}

func main() {
	fmt.Println(recipes([]string{"bread", "sandwich"}, [][]string{{"yeast", "flour"}, {"bread", "meat"}},
		[]string{"yeast", "flour", "meat"}))
}
