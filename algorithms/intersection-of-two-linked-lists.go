package main

import (
	"fmt"
)

// ListNode is the node type for linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pA, pB := headA, headB
	endA, endB := &ListNode{}, &ListNode{}
	lenDiff := 0
	for pA != nil {
		endA = pA
		lenDiff++
		pA = pA.Next
	}
	for pB != nil {
		endB = pB
		lenDiff--
		pB = pB.Next
	}
	if endA == endB {
		pA, pB = headA, headB
		if lenDiff > 0 {
			for lenDiff != 0 {
				pA = pA.Next
				lenDiff--
			}
		} else if lenDiff < 0 {
			for lenDiff != 0 {
				pB = pB.Next
				lenDiff++
			}
		}
		for pA != pB {
			pA = pA.Next
			pB = pB.Next
		}
		return pA
	}
	return nil
}

func main() {
	intersec := &ListNode{234, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	test1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, intersec}}}}
	test2 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, intersec}}}}}
	fmt.Printf("Testing\n%p\n%p\n", test1, test2)
	result := getIntersectionNode(test1, test2)
	fmt.Printf("Yields %p\n", result)

	test1 = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	test2 = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	fmt.Printf("Testing\n%p\n%p\n", test1, test2)
	result = getIntersectionNode(test1, test2)
	fmt.Printf("Yields %p\n", result)
}

/*
8
[4,1,8,4,5]
[5,0,1,8,4,5]
2
3
1
[1]
[1]
0
0
2
[0,9,1,2,4]
[3,2,4]
3
1
*/
