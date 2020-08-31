class Solution:
    def inorderTraversal(self, root: TreeNode) -> List[int]:

        ans = []

        def _recur_(node):
            if node:
                _recur_(node.left)
                ans.append(node.val)
                _recur_(node.right)

        _recur_(root)

        return ans
