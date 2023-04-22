package main

func merge_sorted_array(a []int, n int, b []int, m int) {
    i, j := 0, 0
	res := make([]int, n+m)
	k := 0
	for i < n && j < m {
		if a[i] > b[j] {
			res[k] = b[j]
			j++
		} else {
			res[k] = a[i]
			i++
		}
		k++
	}
	for i < n {
		res[k] = a[i]
		i++
		k++
	}
	for j < m {
		res[k] = b[j]
		j++
		k++
	}
    
    for i:=0;i<n+m;i++{
        a[i] = res[i]
    }
}

func main() {
	merge_sorted_array([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}
