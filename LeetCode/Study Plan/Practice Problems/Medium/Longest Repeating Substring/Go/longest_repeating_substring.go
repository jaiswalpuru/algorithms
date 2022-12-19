package main

import (
	"fmt"
	"math"
	"sort"
)

//-------------------brute force---------------------
func longest_repeating_substring(s string) int {
	n := len(s)
	res := 0
	hash_map := make(map[string]int)
	for i := 0; i < n; i++ {
		t := string(s[i])
		for j := i + 1; j < n; j++ {
			t += string(s[j])
			hash_map[t]++
		}
	}
	arr := []string{}
	for k := range hash_map {
		arr = append(arr, k)
	}
	sort.Strings(arr)
	for i := len(arr) - 1; i >= 0; i-- {
		if hash_map[arr[i]] > 1 && res < len(arr[i]) {
			res = len(arr[i])
		}
	}
	return res
}

//-------------------brute force---------------------

//--------------------efficient approach------------- (binary search + hash_set)
func search(l int, n int, s string) int {
	visited := make(map[string]bool)
	for i := 0; i < n-l+1; i++ {
		temp := s[i : i+l]
		if visited[temp] {
			return i
		}
		visited[temp] = true
	}
	return -1
}

func longest_repeating_substring_eff_one(s string) int {
	l, r := 1, len(s)
	for l <= r {
		mid := (l + r) / 2
		if search(mid, len(s), s) != -1 {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l - 1
}

//--------------------efficient approach-------------

//--------------------efficient approach-------------(rabin karp)
func longest_repeating_substring_eff_two(s string) int {
	n := len(s)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = int(s[i]) - int('a')
	}
	mod := int(math.Pow(float64(2), float64(24)))
	a := 26
	l, r := 1, n
	for l <= r {
		mid := (l + r) / 2
		if search_sub(mid, a, mod, n, arr) != -1 {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l - 1
}

func search_sub(l int, a int, mod, n int, arr []int) int {
	h := 0
	for i := 0; i < l; i++ {
		h = (h*a + arr[i]) % mod
	}
	visited := make(map[int]bool)
	visited[h] = true
	const_val := 1
	for i := 1; i <= l; i++ {
		const_val = (const_val * a) % mod
	}
	for i := 1; i < n-l+1; i++ {
		h = (h*a - arr[i-1]*const_val%mod + mod) % mod
		h = (h + arr[i+l-1]) % mod
		if visited[h] {
			return i
		}
		visited[h] = true
	}
	return -1
}

//--------------------efficient approach-------------

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longest_repeating_substring_eff_two("abcd"))
}
