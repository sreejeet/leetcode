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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := l1
	for ; l2 != nil; l1 = l1.Next {
		sum := l1.Val + l2.Val
		l2 = l2.Next
		l1.Val = sum % 10
		if sum > 9 {
			if l1.Next == nil {
				l1.Next = &ListNode{0, nil}
			}
			l1.Next.Val++
		}
		if l1.Next == nil {
			l1.Next = l2
			println("to...")
			break
		}
	}

	for ; l1 != nil && l1.Val > 9; l1 = l1.Next {
		println("next...")
		l1.Val %= 10
		if l1.Next == nil {
			l1.Next = &ListNode{1, nil}
			break
		}
		l1.Next.Val++
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
		str.WriteString(" " + strconv.Itoa(list.Val))
		list = list.Next
	}
	fmt.Println(str.String())
}

func main() {
	test1 := &ListNode{0, nil}
	test2 := &ListNode{7, &ListNode{3, nil}}
	fmt.Printf("Testing: \n")
	printll(test1)
	printll(test2)
	result := addTwoNumbers(test1, test2)
	fmt.Printf("==========\n")
	printll(result)
	println()
}

/*

[2,4,3]
[5,6,4]
[9]
[0]
[0]
[9]
[9,9,9,9]
[1]
[1]
[9,9,9]
[1,0,0,1]
[9,0,0,9]
[9,9,9]
[1]
[1]
[9,9,9]
[0]
[7,3]

*/