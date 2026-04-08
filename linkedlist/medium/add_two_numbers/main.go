package main

import (
	"github.com/boyuu/LeetCodeGo/linkedlist/utils"
)

func main() {
	l1 := utils.MakeLinkedList(9, 9, 9, 9)
	l2 := utils.MakeLinkedList(9, 9, 9, 9, 9, 9)
	utils.PrintLinkedList(addTwoNumbers(l1, l2))
}

type ListNode = utils.ListNode

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
