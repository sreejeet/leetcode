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
	ret = head
	var prevP *ListNode
	for head != nil && head.Next != nil {
		if head.Val == head.Next.Val {
			nextP := head.Next.Next
			for nextP != nil {
				if nextP.Val != head.Val {
					break
				}
				nextP = nextP.Next
			}
			if prevP == nil {
				head = nextP
				ret = nextP
			} else {
				prevP.Next = nextP
				head = prevP.Next
			}
		} else {
			prevP = head
			head = head.Next
		}
	}
	return
}

func main() {
	test := &ListNode{1, &ListNode{1, &ListNode{2, nil}}}
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
