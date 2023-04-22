package main

import "fmt"

func max(a, b int) int {
	if a>b{
		return a
	}
	return b
}

//-------------------------------brute force----------------------------------(TLE)
func is_valid(s []byte) bool {
	stack := []byte{}
	for i:=0;i<len(s);i++{
		if s[i] == '(' {
			stack = append(stack, s[i])
        }else if len(stack)!= 0 && stack[len(stack)-1]=='('{
            stack = stack[:len(stack)-1]
        }else {
            return false
        }
	}
    return len(stack)==0
}

func longest_valid_parantheses_brute_force(s string) int {
	l := 0
	n := len(s)
	for i:=0;i<len(s);i++{
		for j:=i+2;j<=n;j+=2{
			if is_valid([]byte(s[i:j])) {
				l = max(l, j-i)
			}
		}
	}
	return l
}
//-------------------------------brute force----------------------------------

func longest_valid_parantheses(s string) int {
	str := []byte(s)
	n := len(str)
	l :=0 
	stack := []int{-1}
	for i:=0;i<n;i++{
		if str[i] == '(' {
			stack = append(stack, i)
		}else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0{
				stack = append(stack, i)
			}else {
				l = max(l, i-stack[len(stack)-1])
			}
		}
	}
	return l
}

func main() {
	fmt.Println(longest_valid_parantheses("(()"))
}