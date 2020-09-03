class Solution:
    def minDepth(self, root: TreeNode) -> int:

        if not root:
            return 0

        def _re_(node, lvl):

            if not node:
                return None

            if not node.left and not node.right:
                return lvl

            a = _re_(node.left, lvl+1)
            b = _re_(node.right, lvl+1)
            if a and b:
                return min(a,b)
            return a or b


        return _re_(root, 1)



"""
[3,9,20,null,null,15,7]
[3,9,20,4,5,15,7,null,null,9]
[2,null,3,null,4,null]
[1,null,2,null,3,null,4]
[0]
[]
"""