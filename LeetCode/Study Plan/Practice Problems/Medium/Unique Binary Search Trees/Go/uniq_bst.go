package main

import (
	"fmt"
)

//------------------using dp-------------------------
func unique_bst(n int) int { //refer leetcode soln
	g := make([]int, n+1)
	g[0] = 1
	g[1] = 1

	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			g[i] += g[j-1] * g[i-j]
		}
	}
	return g[n]
}

//------------------using dp-------------------------

//------------------------------using catalan numbers--------------------
// cn = (2*(2n+1)/(n+2))*cn
func unique_bst_catalan(n int) int {
	c := 1
	for i := 0; i < n; i++ {
		c = c * 2 * (2*i + 1) / (i + 2)
	}
	return c
}

//------------------------------using catalan numbers--------------------

func main() {
	fmt.Println(unique_bst_catalan(3))
}
