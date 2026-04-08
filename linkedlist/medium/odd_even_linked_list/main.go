package main

import (
	"github.com/boyuu/LeetCodeGo/linkedlist/utils"
)

func main() {
	root := utils.MakeLinkedList(1, 2, 3, 4, 5)
	root = oddEvenList(root)
	utils.PrintLinkedList(root)
}

type ListNode = utils.ListNode

func oddEvenList(head *ListNode) *ListNode {
	odd := &ListNode{}
	even := &ListNode{}
	pOdd, pEven, p := odd, even, head
	judge := 1
	for p != nil {
		switch {
		case judge == 1:
			pOdd.Next = p
			pOdd = pOdd.Next
		case judge == 0:
			pEven.Next = p
			pEven = pEven.Next
		}
		p = p.Next
		judge = (judge + 1) % 2
	}
	pOdd.Next = even.Next
	pEven.Next = nil // very important
	return odd.Next
}
