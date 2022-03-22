package main

import (
	"fmt"
	"math"
)

func kth_symbol_grammar(n, k int) int {
	if n == 1 && k == 1{
		return 0
	}

	mid := int(math.Pow(2, float64(n-1)))/2
	if k<=mid {
		return kth_symbol_grammar(n-1, k)
	}
	res := kth_symbol_grammar(n-1, k-mid)
	if res == 1{
		return 0
	}else{
		return 1
	}
}

func main() {
	fmt.Println(kth_symbol_grammar(2,2))
}