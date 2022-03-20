package main

import "fmt"

func count_collisions(directions string) int {
	n := len(directions)
	d := []byte(directions)
	col := 0
	for i := 0; i < n-1; i++ {
		if d[i] == 'R' {
			if d[i+1] == 'L' || d[i+1] == 'S' {
				if d[i+1] == 'S' {
					col++
				} else {
					col += 2
				}
				d[i] = 'S'
				d[i+1] = 'S'
			} else if d[i+1] == 'R' {
				cnt := 0
				k := i
				for k < n && d[k] == 'R' {
					cnt++
					k++
				}
				if k == n {
					break
				}
				if d[k] == 'L' {
					col += cnt + 1
				} else {
					col += cnt
				}
				d[k] = 'S'
				i = k - 1
			}
		} else if d[i] == 'S' && d[i+1] == 'L' {
			col++
			d[i+1] = 'S'
		}
	}
	return col
}

func main() {
	fmt.Println(count_collisions("RLRSLL"))
}
