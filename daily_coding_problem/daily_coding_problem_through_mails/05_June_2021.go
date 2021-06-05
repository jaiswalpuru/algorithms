/*
	Implement a queue using two stacks. Recall that a queue is a FIFO (first-in, first-out)
	data structure with the following methods: enqueue,
	which inserts an element into the queue, and dequeue, which removes it.

	Methods :
	Either we can make enqueue operation costly or by making the deque operation costly.
*/

package main

import (
	"fmt"
	"os"
)

func enqueue(val int, stack_one, stack_two *[]int) {

	for len(*stack_one) > 0 {
		v := (*stack_one)[0]
		*stack_one = (*stack_one)[1:]
		*stack_two = append(*stack_two, v)
	}

	*stack_one = append(*stack_one, val)
	for len(*stack_two) > 0 {
		v := (*stack_two)[0]
		*stack_two = (*stack_two)[1:]
		*stack_one = append(*stack_one, v)
	}

}

func deque(stack_one *[]int) int {
	if len(*stack_one) <= 0 {
		fmt.Fprintln(os.Stderr, "queue is empty")
		return -1
	}
	val := (*stack_one)[len(*stack_one)-1]
	*stack_one = (*stack_one)[:len(*stack_one)-1]
	return val
}

func main() {
	stack_one := make([]int, 0)
	stack_two := make([]int, 0)

	enqueue(2, &stack_one, &stack_two)

	fmt.Println(deque(&stack_one))

	enqueue(3, &stack_one, &stack_two)
	enqueue(4, &stack_one, &stack_two)
	fmt.Println(deque(&stack_one))
}
