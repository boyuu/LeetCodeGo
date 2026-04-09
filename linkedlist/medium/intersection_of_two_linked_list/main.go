package main

import (
	"fmt"

	"github.com/boyuu/LeetCodeGo/linkedlist/utils"
)

func main() {
	headA := utils.MakeLinkedList(1, 2, 3, 4, 5)
	headB := utils.MakeLinkedList(9, 8, 7, 6, 5)
	node := getIntersectionNode(headA, headB)
	fmt.Println(node)
}

type ListNode = utils.ListNode

type Iterator struct {
	l1, l2   *ListNode
	cur      *ListNode
	isSecond bool
}

func newIterator(l1, l2 *ListNode) *Iterator {
	return &Iterator{
		l1: l1,
		l2: l2,
		cur: &ListNode{
			Next: l1,
		},
	}
}

func (i *Iterator) Next() *ListNode {
	if i.cur == nil {
		return nil
	}
	i.cur = i.cur.Next
	if i.cur == nil && !i.isSecond {
		i.cur = i.l2
		i.isSecond = true
	}
	return i.cur
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	iter1 := newIterator(headA, headB)
	iter2 := newIterator(headB, headA)
	for {
		p1 := iter1.Next()
		p2 := iter2.Next()
		if p1 == p2 {
			return p1
		}
	}
}
