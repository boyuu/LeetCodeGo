package main

import (
	"container/list"
)

func main() {
	root := makeNode(1, 2, 3, 4, 5, 6, 7)
	connect(root)
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	// recursiveConnect(root)
	iterativeConnect(root)
	return root
}

func iterativeConnect(root *Node) {
	if root == nil {
		return
	}
	l := list.New()
	l.PushBack(root)

	for l.Len() != 0 {
		length := l.Len()
		var prevNode *Node
		for i := 0; i < length; i++ {
			node := l.Remove(l.Front()).(*Node)
			if prevNode != nil {
				prevNode.Next = node
			}
			prevNode = node
			if node.Left != nil {
				l.PushBack(node.Left)
			}
			if node.Right != nil {
				l.PushBack(node.Right)
			}
		}
	}
}

func recursiveConnect(root *Node) {
	if root == nil {
		return
	}
	if root.Left != nil {
		root.Left.Next = root.Right
	}
	if root.Right != nil && root.Next != nil {
		root.Right.Next = root.Next.Left
	}
	recursiveConnect(root.Left)
	recursiveConnect(root.Right)
}

func makeNode(nums ...any) *Node {
	if len(nums) == 0 {
		return nil
	}
	for _, num := range nums {
		if num == nil {
			continue
		}
		if _, ok := num.(int); !ok {
			panic("nums element should be either int or nil")
		}
	}
	root := &Node{Val: nums[0].(int)}
	stack := list.New()
	stack.PushBack(root)
	i := 1
	for stack.Len() > 0 {
		node := stack.Remove(stack.Front()).(*Node)
		if i < len(nums) {
			if nums[i] != nil {
				node.Left = &Node{Val: nums[i].(int)}
				stack.PushBack(node.Left)
			}
			i++
		}
		if i < len(nums) {
			if nums[i] != nil {
				node.Right = &Node{Val: nums[i].(int)}
				stack.PushBack(node.Right)
			}
			i++
		}
	}
	return root
}
