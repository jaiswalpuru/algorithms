package main

import "fmt"
import "strconv"

func to_string(a,b int) string {
	return strconv.Itoa(a) + "," + strconv.Itoa(b)
}


func count_lattice_points_in_circle(circles [][]int) int {
	n := len(circles)
	hash_map := make(map[string]string)
	for i:=0;i<n;i++{
		x1,y1,r1 := circles[i][0],circles[i][1],circles[i][2]
		for x:=-r1;x<=r1;x++{
			for y:=-r1;y<=r1;y++{
				if ((x*x)+(y*y)) <= (r1*r1) {
					hash_map[to_string(x1+x,y1+y)]=to_string(x1+x,y1+y)
				}
			}
		}
	}
	return len(hash_map)
}

func main(){
	fmt.Println(count_lattice_points_in_circle([][]int{
		{2,2,2},{3,4,1},
	}))
}