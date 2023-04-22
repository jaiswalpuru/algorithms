package main

import "fmt"

//--------------------------------worst solution one can ever come up with--------------------------------
func sum_4_bf(nums1, nums2, nums3, nums4 []int) int {
	n, m, p, q := len(nums1), len(nums2), len(nums3), len(nums4)
	cnt := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < p; k++ {
				for l := 0; l < q; l++ {
					if nums1[i]+nums2[j]+nums3[k]+nums4[l] == 0 {
						cnt++
					}
				}
			}
		}
	}
	return cnt
}

//--------------------------------worst solution one can ever come up with--------------------------------

//--------------------------------efficient solution--------------------------------
func sum_4_eff(nums1, nums2, nums3, nums4 []int) int {
	cnt := 0
	mp_4 := make(map[int]int)
	n, m, p, q := len(nums1), len(nums2), len(nums3), len(nums4)
	for i := 0; i < q; i++ {
		for j := 0; j < p; j++ {
			mp_4[nums4[i]+nums3[j]]++
		}
	}
	target := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			val := target - (nums1[i] + nums2[j])
			if _, ok := mp_4[val]; ok {
				cnt += mp_4[val]
			}
		}
	}
	return cnt
}

//------------------------------------------------------------------------------------------------

func main() {
	fmt.Println(sum_4_eff([]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}))
}
