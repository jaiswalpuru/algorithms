package main

import "fmt"
import "sort"

type Node struct {
	Val      int
	Children []*Node
}

func New(val int) *Node { return &Node{Val: val, Children: make([]*Node, 0)} }

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}

//------------------------------without height----------------------------------
func diameter_of_n_ary_tree(root *Node) int {
	res := 0
	dfs(root, &res)
	return res
}

func dfs(root *Node, res *int) int {
	max_arr := []int{0,0}
	n := len(root.Children)
	for i := 0; i < n; i++ {
		curr := dfs(root.Children[i], res)
		max_arr = append(max_arr, curr)
	}

	sort.Ints(max_arr)
	*res = max(*res, max_arr[len(max_arr)-1]+max_arr[len(max_arr)-2])
	return 1+max_arr[len(max_arr)-1]
}
//------------------------------without height-----------------------------------

//-----------------------using height--------------------------------
func diameter_of_n_ary_tree_using_height(root *Node) int {
	dia := 0
	height(root, &dia)
	return dia
}

func height(root *Node, dia *int) int {
	if len(root.Children) == 0{
		return 0
	}

	max_ht_1, max_ht_2 := 0,0 
	for i:=0; i<len(root.Children); i++ {
		parent_ht := height(root.Children[i], dia)+1
		if parent_ht > max_ht_1 {
			max_ht_2 = max_ht_1
			max_ht_1 = parent_ht
		}else if parent_ht > max_ht_2{
			max_ht_2 = parent_ht
		}
		dis := max_ht_1+max_ht_2
		*dia = max(*dia, dis)
	}

	return max_ht_1
}

//-----------------------using height--------------------------------

func main() {
	root := New(1)
	root.Children = append(root.Children, New(3))
	root.Children = append(root.Children, New(2))
	root.Children = append(root.Children, New(4))
	root.Children[0].Children = append(root.Children[0].Children, New(5))
	root.Children[0].Children = append(root.Children[0].Children, New(6))
	fmt.Println(diameter_of_n_ary_tree(root))
}
