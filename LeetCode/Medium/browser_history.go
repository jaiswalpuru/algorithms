/*
You have a browser of one tab where you start on the homepage and you can visit another url,
get back in the history number of steps or move forward in the history number of steps.

Implement the BrowserHistory class:

BrowserHistory(string homepage) Initializes the object with the homepage of the browser.
void visit(string url) Visits url from the current page. It clears up all the forward history.
string back(int steps) Move steps back in history. If you can only return x steps in the history and steps > x,
you will return only x steps. Return the current url after moving back in history at most steps.
string forward(int steps) Move steps forward in history. If you can only forward x steps in the history and steps > x,
you will forward only x steps. Return the current url after forwarding in history at most steps.
*/

package main

import (
	"fmt"
	"strconv"
)

type BrowserHistory struct {
	web_site      string
	ind           int
	state         int
	stack         map[int]string
	stack_reverse map[string]int
}

func Constructor(homepage string) BrowserHistory {
	bh := BrowserHistory{web_site: homepage, ind: 0, state: 0, stack: make(map[int]string), stack_reverse: make(map[string]int)}
	bh.stack[bh.ind] = homepage
	bh.stack_reverse[homepage] = bh.ind
	return bh
}

func (this *BrowserHistory) Visit(url string) {
	if _, ok := this.stack_reverse[url]; ok {
		this.ind = this.stack_reverse[url]
		this.web_site = url
	} else {
		this.state = this.ind + 1
		this.ind = this.state
		this.stack[this.state] = url
		this.stack_reverse[url] = this.state
		this.web_site = url
	}
}

func (this *BrowserHistory) Back(steps int) string {
	if this.ind-steps < 0 {
		this.ind = 0
		this.web_site = this.stack[this.ind]
	} else {
		this.ind -= steps
		this.web_site = this.stack[this.ind]
	}
	return this.web_site
}

func (this *BrowserHistory) Forward(steps int) string {
	if this.ind+steps > this.state {
		this.ind = this.state
		this.web_site = this.stack[this.ind]
		return this.web_site
	}
	this.ind += steps
	this.web_site = this.stack[this.ind]
	return this.web_site
}

func Browser(inst []string, site []string) {

	b := Constructor(site[0])

	n := len(inst)

	for i := 1; i < n; i++ {
		//fmt.Println(b.stack)
		switch inst[i] {
		case "visit":
			fmt.Println("visit", site[i])
			b.Visit(site[i])
		case "back":
			j, _ := strconv.Atoi(site[i])
			fmt.Println("back", b.Back(j), j)
		case "forward":
			j, _ := strconv.Atoi(site[i])
			fmt.Println("forward", b.Forward(j), j)
		}
	}
}

func main() {
	Browser(
		[]string{"BrowserHistory", "visit", "visit", "visit", "back", "back", "forward", "visit", "forward", "back", "back"},
		[]string{"leetcode.com", "google.com", "facebook.com", "youtube.com", "1", "1", "1",
			"linkedin.com", "2", "2", "7"})
}
