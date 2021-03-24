package gostlcode

var ()

//SLBSearch -> return the strictly lower bound index using binary search
func SLBSearch(a []int, v, l, h int) int {

	if len(a) == 0 {
		return 0
	} else if v < a[l] {
		return l
	} else if v >= a[h] {
		return h + 1
	} else {
		m := (l + h) / 2
		if a[m] <= v {
			l = m + 1
		} else {
			h = m - 1
		}
		return SLBSearch(a, v, l, h)
	}

}
