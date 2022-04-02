package main

type LockingTree struct {
	hash_map map[int]map[int]int
	parent   []int
}

func Constructor(parent []int) LockingTree {
	hash_map := make(map[int]map[int]int)
	n := len(parent)
	for i := 0; i < n; i++ {
		if _, ok := hash_map[parent[i]]; !ok {
			hash_map[parent[i]] = make(map[int]int)
		}
		hash_map[parent[i]][i] = -2
	}
	return LockingTree{hash_map, parent}
}

func (this *LockingTree) Lock(num int, user int) bool {
	p := this.parent[num]
	if this.hash_map[p][num] == -2 {
		this.hash_map[p][num] = user
		return true
	}
	return false
}

func (this *LockingTree) Unlock(num int, user int) bool {
	p := this.parent[num]
	if this.hash_map[p][num] == user {
		this.hash_map[p][num] = -2
		return true
	}
	return false
}

func (this *LockingTree) Upgrade(num int, user int) bool {
	t := num
	for {
		p := this.parent[t]
		if p == -1 {
			if this.hash_map[p][t] != -2 {
				return false
			}
			break
		}

		if this.hash_map[p][t] != -2 {
			return false
		}

		t = p

	}

	//
	q := []int{num}
	is_des := false
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		if _, ok := this.hash_map[curr]; ok {
			for k := range this.hash_map[curr] {
				q = append(q, k)
				if this.hash_map[curr][k] != -2 {
					is_des = true
				}
				this.hash_map[curr][k] = -2
			}
		}
	}
	if !is_des {
		return false
	}
	p := this.parent[num]
	this.hash_map[p][num] = user
	return true
}

func main() {

}
