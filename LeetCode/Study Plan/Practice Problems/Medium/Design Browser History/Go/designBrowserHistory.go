package main

import "fmt"

type BrowserHistory struct {
	stack   []string
	pointer int
}

func Constructor(homepage string) BrowserHistory {
	st := []string{homepage}
	return BrowserHistory{stack: st, pointer: 1}
}

func (this *BrowserHistory) Visit(url string) {
	this.stack = this.stack[:this.pointer]
	this.stack = append(this.stack, url)
	this.pointer++
	fmt.Println(this.stack, this.pointer)
}

func (this *BrowserHistory) Back(steps int) string {
	if this.pointer <= steps {
		this.pointer = 1
		return this.stack[0]
	}
	this.pointer -= steps
	return this.stack[this.pointer-1]
}

func (this *BrowserHistory) Forward(steps int) string {
	if this.pointer+steps >= len(this.stack) {
		this.pointer = len(this.stack)
		return this.stack[this.pointer-1]
	}
	this.pointer += steps
	return this.stack[this.pointer-1]
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
