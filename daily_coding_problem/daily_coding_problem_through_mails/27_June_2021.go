/*
Suppose you have a multiplication table that is N by N. That is, a 2D array where the value at the
i-th row and j-th column is (i + 1) * (j + 1) (if 0-indexed) or i * j (if 1-indexed).

Given integers N and X, write a function that returns the number of times X appears as a value in
an N by N multiplication table.

For example, given N = 6 and X = 12, you should return 4, since the multiplication table looks like this:

| 1 | 2 | 3 | 4 | 5 | 6 |

| 2 | 4 | 6 | 8 | 10 | 12 |

| 3 | 6 | 9 | 12 | 15 | 18 |

| 4 | 8 | 12 | 16 | 20 | 24 |

| 5 | 10 | 15 | 20 | 25 | 30 |

| 6 | 12 | 18 | 24 | 30 | 36 |

And there are 4 12's in the table.
*/

package main

import "fmt"

func count_val_in_table(table [][]int, val int) int {
	cnt := 0
	n := len(table)

	for i := 0; i < n; i++ {
		low, high := 0, len(table[i])-1
		for low <= high {
			mid := low + (high-low)/2
			if table[i][mid] == val {
				cnt++
				break
			} else if table[i][mid] > val {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}

	return cnt
}

func main() {
	mul_table := [][]int{
		{1, 2, 3, 4, 5, 6},

		{2, 4, 6, 8, 10, 12},

		{3, 6, 9, 12, 15, 18},

		{4, 8, 12, 16, 20, 24},

		{5, 10, 15, 20, 25, 30},

		{6, 12, 18, 24, 30, 36},
	}
	X := 12
	fmt.Println(count_val_in_table(mul_table, X))
}
