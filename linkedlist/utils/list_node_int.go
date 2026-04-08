package utils

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func MakeLinkedList(nums ...int) *ListNode {
	dummy := &ListNode{}
	p := dummy
	for _, n := range nums {
		p.Next = &ListNode{Val: n}
		p = p.Next
	}
	return dummy.Next
}

func PrintLinkedList(root *ListNode) {
	for root != nil {
		fmt.Print(root.Val, " ")
		root = root.Next
	}
	fmt.Println("")
}
