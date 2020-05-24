package main

import "fmt"

// TreeNode is the node type for a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	q := []*TreeNode{root}
	for len(q) != 0 {
		node := q[0]
		q = q[1:]
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
		node.Left, node.Right = node.Right, node.Left
	}
	return root
}

func main() {
	test := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	fmt.Printf("Testing %v \n", test)
	invertTree(test)
	fmt.Printf("Yields %v\n", test)
}

/*
[]
[1]
[1]
[2]
[0,0]
[1,1]
[2,2]
[0,1]
[1,0]
[1,2,3]
[5,5,5,5]
[2,0,2,1,1,0]
[2,2,1,1,0,0]
[4,2,7,1,3,6,9]
*/
