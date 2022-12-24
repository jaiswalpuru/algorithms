package main

import "fmt"

var mod = int(1e9 + 7)

func domino_tromino_tiling(n int) int {
	full, partial := make(map[int]int), make(map[int]int)
	return full_covered(n, &partial, &full)
}

func partial_covered(n int, partial, full *map[int]int) int {
	if v, ok := (*partial)[n]; ok {
		return v
	}
	v := 0
	if n == 2 {
		v = 1
	} else {
		v = (partial_covered(n-1, partial, full))%mod + (full_covered(n-2, partial, full))%mod
	}
	(*partial)[n] = v
	return v
}

func full_covered(n int, partial, full *map[int]int) int {
	if v, ok := (*full)[n]; ok {
		return v
	}
	v := 0
	if n == 1 {
		v = 1
	} else if n == 2 {
		v = 2
	} else {
		v = (full_covered(n-1, partial, full)%mod + full_covered(n-2, partial, full)%mod + 2*partial_covered(n-1, partial, full)%mod) % mod
	}
	(*full)[n] = v
	return v
}

func main() {
	fmt.Println(domino_tromino_tiling(3))
}
