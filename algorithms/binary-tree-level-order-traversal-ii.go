package main

import "fmt"

// TreeNode is the node type for a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) (res [][]int) {
	res = [][]int{}
	if root == nil {
		return
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		level := []int{}
		for _, x := range q {
			q = q[1:]
			if x.Left != nil {
				q = append(q, x.Left)
			}
			if x.Right != nil {
				q = append(q, x.Right)
			}
			level = append(level, x.Val)
		}
		res = append(res, level)
	}
	for l, r := 0, len(res)-1; l < r; {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return
}

func main() {
	test := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	println("Testing", test)
	result := levelOrderBottom(test)
	fmt.Println("Yields %v\n", result)
}

/*
[3,9,20,null,null,15,7]
[1,2,20,34,34,15,7]
[]
[3]
[3,9]
[3,9,20]
*/
