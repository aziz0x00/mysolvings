/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    bool isBalanced(TreeNode* root) {
        bool ans = true;
        checkhb(root, ans);
        return ans;
    }
    int checkhb(TreeNode* node, bool& ans) {
        if (node == nullptr) return 0;
        int l = checkhb(node->left, ans);
        int r = checkhb(node->right, ans);
        if (abs(l-r) > 1) {
            ans = false;
            return 0;
        } else {
            return 1 + max(l, r);
        }
    }
};
