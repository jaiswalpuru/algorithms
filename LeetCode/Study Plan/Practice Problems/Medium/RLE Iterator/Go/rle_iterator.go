package main

type RLEIterator struct {
	arr []int
}

func Constructor(encoding []int) RLEIterator {
	return RLEIterator{arr: encoding}
}

func (this *RLEIterator) Next(n int) int {
	if len(this.arr) > 0 {
		for i := 0; i < len(this.arr); i += 2 {
			cnt := this.arr[i]
			val := this.arr[i+1]
			if cnt >= n {
				this.arr[i] -= n
				return val
			} else {
				n -= this.arr[i]
				this.arr[i] = 0
			}
		}
	}
	return -1
}
