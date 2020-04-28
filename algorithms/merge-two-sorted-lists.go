package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ListNode type is each node of a linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		if l1 != nil {
			return l1
		}
		return l2
	}
	if l1.Val > l2.Val {
		l1, l2 = l2, l1
	}
	head := l1
	for {
		if l1.Next == nil {
			l1.Next = l2
			return head
		}
		if l1.Next.Val <= l2.Val {
			l1 = l1.Next
			continue
		}
		l1.Next, l2 = l2, l1.Next
	}
	return head
}

func printll(list *ListNode) {
	if list == nil {
		fmt.Println()
		return
	}
	var str strings.Builder
	str.WriteString(strconv.Itoa(list.Val))
	list = list.Next
	for list != nil {
		str.WriteString(" > " + strconv.Itoa(list.Val))
		list = list.Next
	}
	fmt.Println(str.String())
}

func main() {
	testParam1 := &ListNode{1, &ListNode{2, &ListNode{5, &ListNode{6, nil}}}}
	testParam2 := &ListNode{3, &ListNode{4, &ListNode{7, nil}}}
	result := mergeTwoLists(testParam1, testParam2)
	fmt.Printf("Yields : ")
	printll(result)

	testParam1 = &ListNode{2, nil}
	testParam2 = &ListNode{4, &ListNode{7, &ListNode{7, nil}}}
	result = mergeTwoLists(testParam1, testParam2)
	fmt.Printf("Yields : ")
	printll(result)

	testParam1 = &ListNode{6, nil}
	testParam2 = &ListNode{5, &ListNode{7, &ListNode{8, nil}}}
	result = mergeTwoLists(testParam1, testParam2)
	fmt.Printf("Yields : ")
	printll(result)

	testParam1 = &ListNode{9, nil}
	testParam2 = &ListNode{4, &ListNode{7, &ListNode{7, nil}}}
	result = mergeTwoLists(testParam1, testParam2)
	fmt.Printf("Yields : ")
	printll(result)

	testParam1 = nil
	testParam2 = nil
	result = mergeTwoLists(testParam1, testParam2)
	fmt.Printf("Yields : ")
	printll(result)
}
