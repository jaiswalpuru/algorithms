package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//-------------------------using bfs---------------------------------
func longest_path_with_different_adjacent_characters_bfs(parent []int, s string) int {
	n := len(parent)
	childrens := make([]int, n)
	// for each node count the number of childrens except 0
	for i := 1; i < n; i++ {
		childrens[parent[i]]++
	}
	q := []int{}
	//we need to push all the leaf nodes as the path will be build  from
	//bottom to top
	longest_path := 1 //for a single node
	longest_chain := make([][2]int, len(parent))
	for i := 1; i < n; i++ {
		if childrens[i] == 0 { //if zero children that means its the leaf
			longest_chain[i][0] = 1 //each node itself will be of length 1
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		par := parent[curr]
		longest_chain_from_this_node := longest_chain[curr][0]
		if s[curr] != s[par] {
			if longest_chain_from_this_node > longest_chain[par][0] {
				longest_chain[par][1] = longest_chain[par][0]
				longest_chain[par][0] = longest_chain_from_this_node
			} else if longest_chain_from_this_node > longest_chain[par][1] {
				longest_chain[par][1] = longest_chain_from_this_node
			}
		}
		longest_path = max(longest_path, longest_chain[par][0]+longest_chain[par][1]+1)
		childrens[par]-- // reduce the number of children for this parent as it's child has been processed
		if childrens[par] == 0 && par != 0 {
			//if the number of children has been zero for this parent then
			//push it to the queue as it became a child
			q = append(q, par)
			longest_chain[par][0]++
		}
	}
	return longest_path
}

//-------------------------using bfs---------------------------------

//-------------------------using dfs---------------------------------
func dfs(node int, hash_map map[int][]int, res *int, s string) int {
	two_max := []int{0, 0}
	for i := 0; i < len(hash_map[node]); i++ {
		curr := dfs(hash_map[node][i], hash_map, res, s)
		if s[node] != s[hash_map[node][i]] {
			if curr > two_max[0] {
				two_max[1] = two_max[0]
				two_max[0] = curr
			} else if curr > two_max[1] {
				two_max[1] = curr
			}
		}
	}
	*res = max(*res, two_max[0]+two_max[1]+1)
	return two_max[0] + 1
}

func longest_path_with_different_adjacent_characters(parent []int, s string) int {
	hash_map := make(map[int][]int)
	for i := 0; i < len(parent); i++ {
		if parent[i] >= 0 {
			hash_map[parent[i]] = append(hash_map[parent[i]], i)
		}
	}
	res := 0
	dfs(0, hash_map, &res, s)
	return res
}

//-------------------------using dfs---------------------------------

func main() {
	fmt.Println(longest_path_with_different_adjacent_characters([]int{-1, 0, 0, 1, 1, 2}, "abacbe"))
}
