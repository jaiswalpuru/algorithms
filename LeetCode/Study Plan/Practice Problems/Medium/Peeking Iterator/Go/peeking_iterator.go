package main

/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *       
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

type PeekingIterator struct {
    arr []int
    i int
    iter *Iterator
}

func Constructor(iter *Iterator) *PeekingIterator {
    return &PeekingIterator{iter : iter}
}

func (this *PeekingIterator) hasNext() bool {
    return this.iter.index < len(this.iter.v)
}

func (this *PeekingIterator) next() int {
    val := this.iter.v[this.iter.index]
    this.iter.index+=1
    return val
}

func (this *PeekingIterator) peek() int {
    return this.iter.v[this.iter.index]
}

func main(){}