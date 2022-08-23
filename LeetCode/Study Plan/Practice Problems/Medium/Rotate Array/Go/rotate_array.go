package main

func rotate_array(arr []int, k int) {
	n := len(arr)
	if n < k {
		k = k % n
	}
	res := make([]int, n)
	t := 0
	for i := n - k; i < n; i++ {
		res[t] = arr[i]
		t++
	}
	for i := 0; i < n-k; i++ {
		res[t] = arr[i]
		t++
	}
	for i := 0; i < n; i++ {
		arr[i] = res[i]
	}
}

func main() {
	rotate_array([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}
