package main

import (
	"fmt"
	"sort"
)

func smallest_number(n int64) int64 {
	nums := []int{}
	f := false
	if n < 0 {
		f = true
		n = n * -1
	}

	for n > 0 {
		v := int(n) % 10
		nums = append(nums, v)
		n = n / 10
	}

	sort.Ints(nums)

	m := len(nums)

	number := 0
	ten := 10

	if f {
		for i := m - 1; i >= 0; i-- {
			number = (ten * number) + nums[i]
		}
		number *= -1
	} else {
		if len(nums) > 1 && nums[0] == 0 {
			i := 0
			for i < len(nums) && nums[i] == 0 {
				i++
			}
			nums[0], nums[i] = nums[i], nums[0]
		}
		for i := 0; i < m; i++ {
			number = (ten * number) + nums[i]
		}
	}

	return int64(number)
}

func main() {
	fmt.Println(smallest_number(-7605))
}
