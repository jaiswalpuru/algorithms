package main

import "fmt"

func asteroidCollision(asteroids []int) []int {
	size := len(asteroids)
	stack := []int{}

	for i := 0; i < size; i++ {
		if len(stack) == 0 {
			stack = append(stack, asteroids[i])
			continue
		}

		if asteroids[i] < 0 && stack[len(stack)-1] > 0 {
			push := false
			for len(stack) > 0 && (asteroids[i] < 0 && stack[len(stack)-1] > 0) {
				if abs(asteroids[i]) > stack[len(stack)-1] {
					stack = stack[:len(stack)-1]
					push = true
				} else if abs(asteroids[i]) < stack[len(stack)-1] {
					push = false
					break
				} else {
					push = false
					stack = stack[:len(stack)-1]
					break
				}
			}
			if push {
				stack = append(stack, asteroids[i])
			}
		} else {
			stack = append(stack, asteroids[i])
		}
	}

	return stack
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -1 * a
}

func main() {
	fmt.Println(asteroidCollision([]int{5, 10, -5}))
}
