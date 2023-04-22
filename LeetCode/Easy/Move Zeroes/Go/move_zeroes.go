package main

func move_zeroes(arr []int) {
	t := 0
	n := len(arr)
	for i := 0; i < n; i++ {
		if arr[i] != 0 {
			arr[t] = arr[i]
			t++
		}
	}
	for i := t; i < n; i++ {
		arr[i] = 0
	}
}

func main() {
	move_zeroes([]int{0, 1, 0, 3, 12})
}
