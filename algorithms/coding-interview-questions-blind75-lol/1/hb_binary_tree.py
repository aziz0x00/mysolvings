# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def isBalanced(self, root: Optional[TreeNode]) -> bool:
        ans = True
        
        def hb_check(node):
            nonlocal ans
            if node == None: return 0
            l = hb_check(node.left)
            r = hb_check(node.right)
            if abs(l - r) > 1:
                ans = False
                return 0

            return 1 + max(l, r)

        hb_check(root)
        return ans
