package main

type Node struct {
	Val      int
	Children []*Node
}

func New(val int) *Node { return &Node{Val: val, Children: nil} }

func find_root(tree []*Node) *Node {
	if tree == nil {
		return nil
	}

	hash_map := make(map[*Node][]*Node)
	indegree := make(map[*Node]int)
	n := len(tree)

	for i := 0; i < n; i++ {

		if _, ok := hash_map[tree[i]]; !ok {
			hash_map[tree[i]] = make([]*Node, 0)
			indegree[tree[i]] = 0
		}

		for _, v := range tree[i].Children {
			if _, ok := hash_map[v]; !ok {
				hash_map[v] = make([]*Node, 0)
			}
			indegree[v]++
			hash_map[tree[i]] = append(hash_map[tree[i]], v)
		}

	}

	for k, v := range indegree {
		if v == 0 {
			return k
		}
	}
	return nil
}

func main() {

}
