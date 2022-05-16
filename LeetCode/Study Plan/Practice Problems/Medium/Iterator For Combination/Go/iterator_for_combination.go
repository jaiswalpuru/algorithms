package main

type CombinationIterator struct {
    i int
    s []string
}


func Constructor(characters string, combinationLength int) CombinationIterator {
    res := []string{}
    _combination(characters, combinationLength, 0, &res, "")
    sort.Strings(res)
    return CombinationIterator{s:res, i:0}
}

func _combination(s string, k int, ind int, res *[]string, set string) {
    if len(set) == k {
        *res = append(*res, set)
        return 
    }
    
    for i:=ind;i<len(s);i++{
        temp := set + string(s[i])
        _combination(s, k, i+1, res, temp)
        temp = temp[:len(temp)-1]
    }
}

func (this *CombinationIterator) Next() string {
    s := this.s[this.i]
    this.i += 1
    return s
}


func (this *CombinationIterator) HasNext() bool {
    if this.i == len(this.s) {
        return false
    }
    return true
}


/**
 * Your CombinationIterator object will be instantiated and called as such:
 * obj := Constructor(characters, combinationLength);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */