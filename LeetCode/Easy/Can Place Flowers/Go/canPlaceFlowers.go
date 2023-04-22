package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	size := len(flowerbed)

	for i := 0; i < size; {
		if n == 0 {
			break
		}
		if flowerbed[i] == 1 {
			i += 2
		} else {
			if i+1 < size && flowerbed[i+1] == 1 {
				i += 3
			} else {
				flowerbed[i] = 1
				i += 2
				n--
			}
		}
	}

	return n == 0
}

func main() {
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
}
