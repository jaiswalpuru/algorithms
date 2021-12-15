// This can be done in linear time , pop until you find a maxium value then push back all the elements

// To get the maximum value in constant time we can maintain two stack one normal and other maxStack which will
// contain maximum values for the stack, it will have the same number of elements as present in original stack
// only it will contain the maximum value at that time.

package main

import "log"

type stack []int

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Pop(st, maxStack *stack) int {
	if len(*st) == 0 {
		return -1
	}
	*maxStack = (*maxStack)[:len(*maxStack)-1]
	val := (*st)[len(*st)-1]
	*st = (*st)[:len(*st)-1]
	return val
}

func Push(st, maxStack *stack, val int) {
	*st = append(*st, val)
	if len(*maxStack) > 0 {
		*maxStack = append(*maxStack, max(val, (*maxStack)[len(*maxStack)-1]))
	} else {
		*maxStack = append(*maxStack, val)
	}
}

func MaxVal(maxStack stack) int {
	if len(maxStack) == 0 {
		return -1
	}
	return maxStack[len(maxStack)-1]
}

func main() {
	st := make(stack, 0)
	maxStack := make(stack, 0)
	t := Pop(&st, &maxStack)
	if t == -1 {
		log.Println("Empty stack")
	} else {
		log.Println("Poppped Val : ", t)
	}
	Push(&st, &maxStack, 12)
	Push(&st, &maxStack, 10)
	Push(&st, &maxStack, 9)
	Push(&st, &maxStack, 45)
	Push(&st, &maxStack, 55)
	Push(&st, &maxStack, 1)
	Push(&st, &maxStack, 2)
	t = MaxVal(maxStack)
	if t == -1 {
		log.Println("Empty stack")
	} else {
		log.Println("Maximum Value ", t)
	}
}
