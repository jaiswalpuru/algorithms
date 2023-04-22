package main

import "fmt"

//--------------------------brute force-------------------------
func longest_duplicate_substring(s string) string {
	mp := make(map[string]int)
	n := len(s)
	for i := 0; i < n; i++ {
		str := string(s[i])
		mp[str]++
		for j := i + 1; j < n; j++ {
			str += string(s[j])
			mp[str]++
		}
	}
	res := ""
	for k, v := range mp {
		if v >= 2 {
			if len(k) > len(res) {
				res = k
			}
		}
	}
	return res
}

//--------------------------brute force-------------------------

//------------------efficient approach---------------------
func longest_duplicate_substring_eff(s string) string {
	n := len(s)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = int(s[i]) - int('a')
	}
	a := 26
	mod := int(1e9 + 7)
	l, r := 1, n
	for l <= r {
		mid := (l + r) / 2
		if search(s, arr, a, mod, mid, n) != -1 {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	start := search(s, arr, a, mod, l-1, n)
	return s[start : start+l-1]
}

func search(s string, arr []int, a, mod, mid, n int) int {
	h := 0
	for i := 0; i < mid; i++ {
		h = (h*a + arr[i]) % mod
	}
	vis := make(map[int][]int)
	vis[h] = append(vis[h], 0)
	const_val := 1
	for i := 1; i <= mid; i++ {
		const_val = (const_val * a) % mod
	}
	for i := 1; i < n-mid+1; i++ {
		h = (h*a - arr[i-1]*const_val%mod + mod) % mod
		h = (h + arr[i+mid-1]) % mod
		hits := vis[h]
		if hits != nil {
			curr := s[i : i+mid]
			for k := 0; k < len(hits); k++ {
				cand := s[hits[k] : hits[k]+mid]
				if cand == curr {
					return hits[k]
				}
			}
		}
		vis[h] = append(vis[h], i)
	}
	return -1
}

//------------------efficient approach---------------------

func main() {
	fmt.Println(longest_duplicate_substring_eff("banana"))
}
