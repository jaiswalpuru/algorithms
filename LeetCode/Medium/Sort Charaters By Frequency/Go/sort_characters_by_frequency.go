package main

import "fmt"
import "container/heap"

type Pair struct {
    char byte
    val int
}

type P []Pair
func (p P) Len() int { return len(p) }
func (p P) Swap(i,j int) { p[i], p[j] = p[j], p[i] }
func (p P) Less(i,j int) bool { return p[i].val > p[j].val }
func (p *P) Push(val interface{}){ *p = append(*p, val.(Pair)) }
func (p *P) Pop() interface {} {
    val := (*p)[len(*p)-1]
    *p = (*p)[:len(*p)-1]
    return val
}


func sort_char_by_frequency(s string) string {
    freq := make(map[byte]int)
    n := len(s)
    str := []byte(s)
    for i:=0;i<n;i++{
        freq[str[i]]++
    }

    var mh P
    for k,v := range freq {
       mh = append(mh, Pair{k,v})
    }
    
    heap.Init(&mh)
    res := ""
    for mh.Len()>0 {
        curr := heap.Pop(&mh).(Pair)
        for i:=0;i<curr.val;i++{
            res += string(curr.char)
        }
    }
    return res
}

func main() {
    fmt.Println(sort_char_by_frequency("tree"))
}
