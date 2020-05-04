package main

import (
	// `golint` keeps adding imports used in other .go files
	// Please ignore if they are not actually used here
	"fmt"
	"strconv"
	"strings"
)

// ListNode is the node type for linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) (ret *ListNode) {
	if head == nil {
		return
	}
	ret = head
	for head.Next != nil {
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}
	return
}

func main() {
	test := &ListNode{1, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}
	test.Print()
	result := deleteDuplicates(test)
	result.Print()
}

// Print is a util func to print linked lists
func (list *ListNode) Print() {
	if list == nil {
		fmt.Println()
		return
	}
	var str strings.Builder
	str.WriteString(strconv.Itoa(list.Val))
	list = list.Next
	for list != nil {
		str.WriteString(" " + strconv.Itoa(list.Val))
		list = list.Next
	}
	fmt.Println(str.String())
}
