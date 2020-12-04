package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func printf(f string, args ...interface{}) { fmt.Fprintf(writer, f, args...) }
func scanf(f string, args ...interface{})  { fmt.Fscanf(reader, f, args...) }

func pref(c rune) int {

	if c == '^' {
		return 3
	} else if c == '*' || c == '/' {
		return 2
	} else if c == '+' || c == '-' {
		return 1
	}
	return -1

}

func main() {

	defer writer.Flush()

	var t int
	scanf("%d\n", &t)

	for ; t > 0; t-- {

		var (
			n   int
			p   = -1
			str string
		)

		scanf("%d\n", &n)
		scanf("%s\n", &str)

		arr := []rune(str)
		stack := make([]rune, len(str))
		postFix := []rune{}

		for i := 0; i < n; i++ {

			if arr[i] >= 'A' && arr[i] <= 'Z' {
				postFix = append(postFix, arr[i])
			} else if arr[i] == ')' {
				for stack[p] != '(' {
					postFix = append(postFix, stack[p])
					stack[p] = ' '
					p--
				}
				stack[p] = ' '
				p--
			} else if arr[i] == '(' {
				p++
				stack[p] = arr[i]
			} else {
				for p != -1 && pref(arr[i]) <= pref(stack[p]) {
					postFix = append(postFix, stack[p])
					stack[p] = ' '
					p--
				}
				p++
				stack[p] = arr[i]
			}
		}

		for p != -1 {
			if stack[p] == '(' {
				stack[p] = ' '
			} else {
				postFix = append(postFix, stack[p])
				stack[p] = ' '
			}
			p--
		}

		printf("%s\n", (string(postFix)))

	}

}
