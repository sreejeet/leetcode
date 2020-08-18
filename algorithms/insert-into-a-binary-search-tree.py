class Solution:
    def insertIntoBST(self, root: TreeNode, val: int) -> TreeNode:
        if not root:
            return TreeNode(val)

        node = root
        while True:
            if val > node.val:
                if node.right != None:
                    node = node.right
                    continue
                node.right = TreeNode(val)
                break
            else:
                if node.left != None:
                    node = node.left
                    continue
                node.left = TreeNode(val)
                break

        return root
