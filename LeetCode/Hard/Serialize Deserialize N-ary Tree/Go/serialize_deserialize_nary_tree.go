package main

import (
	"strconv"
	"strings"
)

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Codec struct{}

func Constructor() *Codec { return &Codec{} }

func (this *Codec) serialize(root *Node) string {
	if root == nil {
		return ""
	}
	str := strconv.Itoa(root.Val) + ",null,"
	q := []*Node{root}
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			parent := q[0]
			q = q[1:]
			if len(parent.Children) == 0 {
				str += "null,"
				continue
			}
			for j := 0; j < len(parent.Children); j++ {
				str += strconv.Itoa(parent.Children[j].Val) + ","
				q = append(q, parent.Children[j])
			}
			str += "null,"
		}
	}
	str = str[:len(str)-1]
	data := strings.Split(str, ",")
	i := 0
	for i = len(data) - 1; i >= 0; i-- {
		if data[i] != "null" {
			break
		}
	}
	data = data[:i+1]
	str = strings.Join(data, ",")
	return str
}

func New(val int) *Node { return &Node{Val: val} }

func (this *Codec) deserialize(data string) *Node {
	if data == "" {
		return nil
	}
	i := 0
	nodes := strings.Split(data, ",")
	n := len(nodes)
	val, _ := strconv.Atoi(nodes[i])
	root := New(val)
	q := []*Node{root}
	i += 2
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		if i >= n {
			break
		}
		if nodes[i] == "null" {
			i++
			continue
		}
		for i < n && nodes[i] != "null" {
			val, _ = strconv.Atoi(nodes[i])
			child := New(val)
			curr.Children = append(curr.Children, child)
			q = append(q, child)
			i++
		}
		i++
	}
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
