package main

import (
	"fmt"
)

// ListNode is the node type for linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	mid, end := head, head
	for end.Next != nil {
		end = end.Next
		if end.Next == nil {
			break
		}
		end, mid = end.Next, mid.Next
	}

	var prePrev *ListNode
	prev := mid
	for prev != nil {
		prev.Next, prePrev, prev = prePrev, prev, prev.Next
	}

	mid.Next = nil
	for head != nil {
		if end.Val != head.Val {
			return false
		}
		end, head = end.Next, head.Next
	}
	return true
}

func main() {
	test := &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{1, nil}}}}
	fmt.Printf("Testing %v \n", test)
	result := isPalindrome(test)
	fmt.Printf("Yields %v\n", result)
}

/*
[]
[1]
[1,2]
[1,2,1,2]
[1,2,2,1]
[1,1,1,1]
[1,2,3,2,1]
[1,1,2,1]
[1,2,2,1,5]
[5,1,2,2,1,7]
[7,1,2,2,1]
*/
