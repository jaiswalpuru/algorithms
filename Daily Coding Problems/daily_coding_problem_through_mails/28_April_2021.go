/*
Reservoir sampling is a family of randomized algorithms for randomly choosing k samples from a list of n items,
where n is either a very large number or unknown number.
Typically n is very large that doesn't fit in main memory. eg : a list of queries in google and facebook.
So we are given a big array of numbers, and we need to write an efficient function to randomly select k numbers
where 1 ≤ k ≤ n. Let the input array be stream[].
A simple soln is to create a arrary reservoir[] of max k. One by one randomly select an item from stream[0..n-1]
If the selected item is not previously selected then put it in the reservoir. To select if the item was previously
selected or not we need to search in the reservoir list. This approach will take O(k^2).
Efficient soln to this is O(n):
1. Create an array reservoir[0...k-1] and copy the first k items of stream[] to it.
2. One by one consider all the items from k+1th items to the nth item.
	a) Generate random number from 0 to i where i is the index of current item in stream[].
	Let the genereted number be j.
	b) if j is in the range of 0 to k-1, replace reservoir[j] with stream[i].
*/

// Given a stream of elements too large to store in memory,
// pick a random element from the stream with uniform probability.
/*
This problem is a variation of reservoir sampling with value of k = 1
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	stream := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	cnt := 0 //denotes the numbers seen so far
	var res int

	for i := 0; i < len(stream); i++ {
		cnt++
		if cnt == 1 {
			res = stream[i]
		} else {
			t := rand.Int() % cnt
			if t == cnt-1 {
				res = stream[i]
			}
		}
		fmt.Printf("Random number for the first %d numbers is %d\n", i+1, res)
	}

}
