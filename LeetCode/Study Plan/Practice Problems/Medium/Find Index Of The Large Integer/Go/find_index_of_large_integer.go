package main

/**
 * // This is the ArrayReader's API interface.
 * // You should not implement it, or speculate about its implementation
 * type ArrayReader struct {
 * }
 * // Compares the sum of arr[l..r] with the sum of arr[x..y]
 * // return 1 if sum(arr[l..r]) > sum(arr[x..y])
 * // return 0 if sum(arr[l..r]) == sum(arr[x..y])
 * // return -1 if sum(arr[l..r]) < sum(arr[x..y])
 * func (this *ArrayReader) compareSub(l, r, x, y int) int {}
 *
 * // Returns the length of the array
 * func (this *ArrayReader) length() int {}
 */

func getIndex(reader *ArrayReader) int {
	l, r := 0, reader.length()-1
	for l < r {
		mid := l + (r-l)/2
		if (r-l)%2 == 0 { //check even elements
			res := reader.compareSub(l, mid-1, mid+1, r)
			if res == 0 {
				l = mid
			} else if res == 1 {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			res := reader.compareSub(l, mid, mid+1, r)
			if res == 1 {
				r = mid
			} else {
				l = mid + 1
			}
		}
	}
	return l
}
