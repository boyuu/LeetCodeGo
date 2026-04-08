package main

import "fmt"

func main() {
	l1 := makeLinkedList(9, 9, 9, 9)
	l2 := makeLinkedList(9, 9, 9, 9, 9, 9)
	printLinkedList(addTwoNumbers(l1, l2))
}

func makeLinkedList(nums ...int) *ListNode {
	dummy := &ListNode{}
	p := dummy
	for _, n := range nums {
		p.Next = &ListNode{Val: n}
		p = p.Next
	}
	return dummy.Next
}

func printLinkedList(root *ListNode) {
	for root != nil {
		fmt.Print(root.Val, " ")
		root = root.Next
	}
	fmt.Println("")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	p := dummy
	remain := 0
	for l1 != nil || l2 != nil {
		sum := remain
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		p.Next = &ListNode{
			Val: sum % 10,
		}
		remain = sum / 10
		p = p.Next
	}
	if remain != 0 {
		p.Next = &ListNode{
			Val: remain,
		}
	}
	return dummy.Next
}
