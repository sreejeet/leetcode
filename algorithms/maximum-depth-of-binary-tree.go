package main

// TreeNode is the node type for a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root != nil {
		dep1 := maxDepth(root.Left)
		dep2 := maxDepth(root.Right)
		if dep1 > dep2 {
			return 1 + dep1
		}
		return 1 + dep2
	}
	return 0
}

func main() {
	test := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	println("Testing")
	result := maxDepth(test)
	println("Yields", result)
}
