type MyQueue struct {
	st_one []int
	st_two []int
}

func Constructor() MyQueue {
	return MyQueue{st_one: make([]int, 0), st_two: make([]int, 0)}
}

func (this *MyQueue) Push(x int) {
	for len(this.st_two) > 0 {
		this.st_one = append(this.st_one, this.st_two[len(this.st_two)-1])
		this.st_two = this.st_two[:len(this.st_two)-1]
	}
	this.st_one = append(this.st_one, x)
	for len(this.st_one) > 0 {
		this.st_two = append(this.st_two, this.st_one[len(this.st_one)-1])
		this.st_one = this.st_one[:len(this.st_one)-1]
	}
}

func (this *MyQueue) Pop() int {
	val := this.st_two[len(this.st_two)-1]
	this.st_two = this.st_two[:len(this.st_two)-1]
	return val
}

func (this *MyQueue) Peek() int {
	if len(this.st_two) != 0 {
		return this.st_two[len(this.st_two)-1]
	}
	return this.st_one[len(this.st_one)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.st_one) == 0 && len(this.st_two) == 0
}

/**
* Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
*/
