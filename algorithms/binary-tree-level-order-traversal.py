class Solution:
    def levelOrder(self, root: TreeNode) -> List[List[int]]:
        
        if not root:
            return []
        
        ans = []
        q = [(root, 0)]

        while q:
            cur_node,cur_lvl = q.pop()

            while len(ans) < cur_lvl+1:
                ans.append([])

            ans[cur_lvl].append(cur_node.val)

            if cur_node.right:
                q.append((cur_node.right, cur_lvl+1))
            if cur_node.left:
                q.append((cur_node.left, cur_lvl+1))

        return ans
