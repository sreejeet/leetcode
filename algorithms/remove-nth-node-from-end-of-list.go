package main

import (
	"fmt"
)

// ListNode is the node type for linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) (first *ListNode) {
	first = head
	var r *ListNode
	for ; head != nil; head = head.Next {
		if n <= 0 {
			if r == nil {
				r = first
				continue
			}
			r = r.Next
			continue
		}
		n--
	}
	if r == nil {
		return first.Next
	}
	if r.Next == nil {
		r.Next = nil
		return
	}
	r.Next = r.Next.Next
	return
}

func main() {
	test1 := &ListNode{1, nil}
	test2 := 1
	// fmt.Printf("Testing %v \n", test1)
	result := removeNthFromEnd(test1, test2)
	fmt.Printf("Yields %v\n", result)

	test1 = &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	test2 = 1
	// fmt.Printf("Testing %v \n", test1)
	result = removeNthFromEnd(test1, test2)
	fmt.Printf("Yields %v\n", result)

	test1 = &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	test2 = 2
	// fmt.Printf("Testing %v \n", test1)
	result = removeNthFromEnd(test1, test2)
	fmt.Printf("Yields %v\n", result)

	test1 = &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	test2 = 3
	// fmt.Printf("Testing %v \n", test1)
	result = removeNthFromEnd(test1, test2)
	fmt.Printf("Yields %v\n", result)
}

/*
[1]
1
[1,2,3]
1
[1,2,3]
2
[1,2,3]
3
*/
