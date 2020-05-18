package main

import (
	"fmt"
)

// TreeNode is the node type for a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) (trav [][]int) {
	trav = [][]int{}
	if root == nil {
		return
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		level := []int{}
		for _, v := range q {
			level = append(level, v.Val)
			if v.Left != nil {
				q = append(q, v.Left)
			}
			if v.Right != nil {
				q = append(q, v.Right)
			}
			q = q[1:]
		}
		trav = append(trav, level)
	}
	return
}

func main() {
	test := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	// test := &TreeNode{1, nil, &TreeNode{2, nil, &TreeNode{3, nil, &TreeNode{4, nil, &TreeNode{5, nil, nil}}}}}
	fmt.Printf("Testing %v \n", test)
	result := levelOrder(test)
	fmt.Printf("Yields %v\n", result)
}

/*
[3,9,20,null,null,15,7]
[1,2,3,4,5,6,7,8,9,0]
[]
[1]
[1,2]
[1,3,4]
[1,nil,nil]
*/
