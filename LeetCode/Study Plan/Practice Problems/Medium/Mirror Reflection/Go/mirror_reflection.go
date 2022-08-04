package main

import "fmt"

func mirror_reflection(p, q int) int {
	x, y := 0, 0
	dir := 1
	for {
		if dir == 1 {
			y += q
			if y > p {
				dir = -1
				y = 2*p - y
			}
		} else {
			y -= q
			if y < 0 {
				dir = 1
				y = -y
			}
		}

		if x == p {
			x = 0
		} else {
			x = p
		}

		if x == p && y == 0 {
			return 0
		} else if x == p && y == p {
			return 1
		} else if x == 0 && y == p {
			return 2
		}
	}

}

func main() {
	fmt.Println(mirror_reflection(3, 2))
}
