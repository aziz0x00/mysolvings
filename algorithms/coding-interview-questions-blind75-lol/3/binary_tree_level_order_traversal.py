# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def levelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        if root == None:
            return []
        q = [root]
        ans = []
        while q:
            n = len(q)
            ans.append([])
            for i in range(n):
                p = q.pop(0)
                ans[-1].append(p.val)
                if p.left:
                    q.append(p.left)
                if p.right:
                    q.append(p.right)
        return ans
