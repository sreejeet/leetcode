package main

import "fmt"

// ListNode is the node type for linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	for p1, p2 := head, head.Next; p1 != p2; p1 = p1.Next {
		if p2 == nil || p2.Next == nil {
			return false
		}
		p2 = p2.Next.Next
	}
	return true
}

func main() {
	test := &ListNode{1, &ListNode{1, nil}}
	test.Next.Next = test
	fmt.Println("Testing", test)
	result := hasCycle(test)
	fmt.Printf("Yields %v\n", result)

	test = &ListNode{1, &ListNode{1, nil}}
	fmt.Println("Testing", test)
	result = hasCycle(test)
	fmt.Printf("Yields %v\n", result)
}

/*
[1,2]
-1
[1]
0
[1,2,3]
0
[1,2,3,4]
3
[1]
-1
[]
-1
*/
