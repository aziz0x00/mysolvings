/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */

class Solution {
public:
    TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
        TreeNode* seek = root;
        for(;;) {
            if (seek->val > p->val && seek->val > q->val) {
                seek = seek->left;
            }
            else if (seek->val < p->val && seek->val < q->val) {
                seek = seek->right;
            }
            else {
                return seek;
            }
            
        }
    }
};
