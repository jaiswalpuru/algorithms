package main

import "fmt"

func partition_labels(str string) []int {
	res := []int{}
	n := len(str)
	
	index_map := make(map[byte]int)
	for i:=0; i<n; i++ {
		index_map[str[i]] = i
	}

	prev := 0
	ind := index_map[str[0]]
	for i:=0; i<n; i++ {
		if i == ind {
			res = append(res, ind-prev+1)
			if i!=n-1{
				prev = ind+1
				ind = index_map[str[i+1]]
			}
		}else {
			ind = max(ind, index_map[str[i]])
		}
	}

	return res
}

func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}

func main() {
	fmt.Println(partition_labels("ababcbacadefegdehijhklij"))
}