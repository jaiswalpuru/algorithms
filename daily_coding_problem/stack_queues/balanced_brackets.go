package main

import "fmt"

type stack []rune

func (st *stack) Pop() rune {
	if len(*st) == 0 {
		return -1
	}
	val := (*st)[len(*st)-1]
	*st = (*st)[:len(*st)-1]
	return val
}

func (st *stack) Empty() bool {
	if len(*st) == 0 {
		return true
	} else {
		return false
	}
}

func (st *stack) Push(val rune) { *st = append(*st, val) }

func (st *stack) Top() rune { return (*st)[len(*st)-1] }

func checkBalancedOrNot(str string) bool {
	st := make(stack, 0)
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		if r[i] == '(' || r[i] == '{' || r[i] == '[' {
			st.Push(r[i])
		} else {
			if (r[i] == ']' && st.Top() == '[') || (r[i] == '}' && st.Top() == '{') || (r[i] == ')' && st.Top() == '(') {
				_ = st.Pop()
			} else {
				return false
			}
		}
	}
	if st.Empty() {
		return true
	}
	return false
}

func main() {
	str1 := "([])[]({})" //balanced brackets
	str2 := "([)]"       //not balanced brackets
	str3 := "((()"       //not balanced
	fmt.Println(str1, checkBalancedOrNot(str1))
	fmt.Println(str2, checkBalancedOrNot(str2))
	fmt.Println(str3, checkBalancedOrNot(str3))
}
