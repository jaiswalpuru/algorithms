package main

import (
	"fmt"
	"sort"
	"strconv"
)

func to_string(a, b, c int) string { return strconv.Itoa(a) + strconv.Itoa(b) + strconv.Itoa(c) }

var mod = 1e9 + 7

//-----------------------------do not use this method--------------------------
func three_sum_with_multiplicity(arr []int, target int) int {
	sort.Ints(arr)
	n := len(arr)

	t := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if arr[i]+arr[j]+arr[k] == target {
					t = (t + 1) % mod
				}
			}
		}
	}
	return t
}

//-----------------------------do not use this method--------------------------

//----------------------------three sum with multiplicity---------------------------
func three_sum_with_multiplicity(arr []int, target int) int {
	sort.Ints(arr)
	n := len(arr)
	hash_map := make(map[int]int)
	uniq := 0
	keys := []int{}

	for i := 0; i < n; i++ {
		hash_map[arr[i]]++
		if hash_map[arr[i]] == 1 {
			keys = append(keys, arr[i])
			uniq++
		}
	}

	ans := 0
	for i := 0; i < len(keys); i++ {
		x := keys[i]
		t := target - x
		j, k := i, len(keys)-1
		for j <= k {
			y, z := keys[j], keys[k]
			if y+z < t {
				j++
			} else if y+z > t {
				k--
			} else {
				if i < j && j < k {
					ans += (hash_map[x] * hash_map[y] * hash_map[z])
				} else if i == j && j < k {
					ans += hash_map[x] * (hash_map[x] - 1) / 2 * hash_map[z]
				} else if i < j && j == k {
					ans += hash_map[x] * hash_map[y] * (hash_map[y] - 1) / 2
				} else { //i==j==k
					ans += (hash_map[x]) * (hash_map[x] - 1) * (hash_map[x] - 2) / 6
				}

				ans %= int(mod)
				j++
				k--
			}
		}
	}
	return ans
}

//----------------------------three sum with multiplicity---------------------------

func main() {
	fmt.Println(three_sum_with_multiplicity([]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}, 8))
}
