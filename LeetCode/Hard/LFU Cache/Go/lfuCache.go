type Node struct {
	key, value, freq int
}

type LFUCache struct {
	nodes    map[int]*list.Element
	lists    map[int]*list.List
	capacity int
	minFreq  int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		nodes:    make(map[int]*list.Element),
		lists:    make(map[int]*list.List),
		capacity: capacity,
		minFreq:  0,
	}
}

func (this *LFUCache) Get(key int) int {
	val, ok := this.nodes[key]
	if !ok {
		return -1
	}
	node := val.Value.(*Node)
	this.lists[node.freq].Remove(val)
	node.freq++
	if _, ok := this.lists[node.freq]; !ok {
		this.lists[node.freq] = list.New()
	}
	freqList := this.lists[node.freq]
	newNode := freqList.PushBack(node)
	this.nodes[key] = newNode
	if node.freq-1 == this.minFreq && this.lists[node.freq-1].Len() == 0 {
		this.minFreq++
	}
	return node.value
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}
	val, ok := this.nodes[key]
	if ok {
		node := val.Value.(*Node)
		node.value = value
		this.Get(key)
		return
	}
	// if the cache size has reached the capacity limit
	if this.capacity == len(this.nodes) {
		freqList := this.lists[this.minFreq]
		front := freqList.Front()
		delete(this.nodes, front.Value.(*Node).key)
		freqList.Remove(front)
	}
	//reset the minimum freq
	this.minFreq = 1
	node := &Node{key: key, value: value, freq: this.minFreq}
	if _, ok := this.lists[this.minFreq]; !ok {
		this.lists[this.minFreq] = list.New()
	}
	freqList := this.lists[this.minFreq]
	newNode := freqList.PushBack(node)
	this.nodes[key] = newNode
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
