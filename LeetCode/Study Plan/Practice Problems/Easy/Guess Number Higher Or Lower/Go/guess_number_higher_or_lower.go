package main

func guess_number(n int) int {
	l, r := 0, n
	for l <= r {
		g := (l + r) / 2
		guess_val := guess(g)
		if guess_val == 0 {
			return g
		}
		if guess_val == -1 {
			r = g - 1
		} else {
			l = g + 1
		}
	}
	return -1
}
