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

func averageOfLevels(root *TreeNode) (avgs []float64) {
	avgs = []float64{}
	if root == nil {
		return
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		average, count := 0, float64(len(q))
		for _, v := range q {
			average += v.Val
			if v.Left != nil {
				q = append(q, v.Left)
			}
			if v.Right != nil {
				q = append(q, v.Right)
			}
			q = q[1:]
		}
		avgs = append(avgs, float64(average)/count)
	}
	return
}

func main() {
	test := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Printf("Testing %v \n", test)
	result := averageOfLevels(test)
	fmt.Printf("Yields %v\n", result)
}

/*
[3,9,20,15,7]
[3,9,20,null,null,15,7]
[1,2,3,4,5,6,7,8,9,0]
[]
[1]
[1,2]
[1,3,4]
[1,nil,nil]
*/
