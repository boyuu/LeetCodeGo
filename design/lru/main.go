package main

import (
	"fmt"
)

func main() {
	solution := Constructor(2)
	solution.set(1, 1)
	solution.set(2, 2)
	fmt.Println(solution.get(1))
	solution.set(3, 3)
	fmt.Println(solution.get(2))
	solution.set(4, 4)
	fmt.Println(solution.get(1))
	fmt.Println(solution.get(3))
	fmt.Println(solution.get(4))
}

type Solution struct {
	// write code here
	KeyMap     map[int]*LinkNode
	Head, Tail *LinkNode
	Capacity   int
	Num        int
}

type LinkNode struct {
	Key, Value int
	Next, Prev *LinkNode
}

func Constructor(capacity int) Solution {
	// write code here
	head := &LinkNode{}
	tail := &LinkNode{}
	head.Next = tail
	tail.Prev = head
	return Solution{
		KeyMap:   make(map[int]*LinkNode),
		Head:     head,
		Tail:     tail,
		Capacity: capacity,
	}
}

func (this *Solution) get(key int) int {
	// write code here
	if node, ok := this.KeyMap[key]; ok {
		this.remove(node)
		this.moveToHead(node)
		return node.Value
	} else {
		return -1
	}
}

func (this *Solution) set(key int, value int) {
	// write code here
	if node, ok := this.KeyMap[key]; ok {
		this.remove(node)
		this.moveToHead(node)
		node.Value = value
		return
	}
	node := &LinkNode{
		Key:   key,
		Value: value,
	}
	this.KeyMap[key] = node
	this.moveToHead(node)
	this.Num++
	if this.Num > this.Capacity {
		delete(this.KeyMap, this.Tail.Prev.Key)
		this.remove(this.Tail.Prev)
	}
}

func (this *Solution) remove(node *LinkNode) {
	prev := node.Prev
	next := node.Next
	prev.Next = next
	next.Prev = prev
}

func (this *Solution) moveToHead(node *LinkNode) {
	next := this.Head.Next
	this.Head.Next = node
	node.Prev = this.Head
	node.Next = next
	next.Prev = node
}
