import sys

class Solution:
    def minDiffInBST(self, root: TreeNode) -> int:

        min_val = [sys.maxsize]

        def _recur_(node):
            if min_val[0] == 1:
                return 1
            if node.left:
                cur = node.left
                while cur.right:
                    cur = cur.right
                diff_l = abs(node.val - cur.val)
                min_val[0] = min(min_val[0], diff_l)
                print(min_val[0])
                _recur_(node.left)
            if node.right:
                cur = node.right
                while cur.left:
                    cur = cur.left
                diff_r = abs(node.val - cur.val)
                min_val[0] = min(min_val[0], diff_r)
                print(min_val[0])
                _recur_(node.right)

        _recur_(root)

        return min_val[0]

