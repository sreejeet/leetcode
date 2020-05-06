package main

// TreeNode is the node type for a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func iCantChangeTheFuncSignature(p *TreeNode, q *TreeNode) bool {
	if p != nil && q != nil {
		return p.Val == q.Val &&
			iCantChangeTheFuncSignature(p.Left, q.Right) &&
			iCantChangeTheFuncSignature(p.Right, q.Left)
	}
	if p == nil && q == nil {
		return true
	}
	return false
}

func isSymmetric(root *TreeNode) bool {
	return iCantChangeTheFuncSignature(root, root)
}

func main() {
	test := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	println("Testing")
	result := isSymmetric(test)
	println("Yields", result)

	test = &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{2, nil, nil}}
	println("Testing")
	result = isSymmetric(test)
	println("Yields", result)
}
