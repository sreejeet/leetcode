package main

import (
	"fmt"
)

// ListNode type is each node of a linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) (prev *ListNode) {
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return
}

func main() {
	test := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	fmt.Printf("Testing %v \n", test)
	result := reverseList(test)
	fmt.Printf("Yields %v\n", result)
}
