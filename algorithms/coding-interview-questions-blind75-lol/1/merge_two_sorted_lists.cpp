/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* mergeTwoLists(ListNode* list1, ListNode* list2) {
        ListNode* head = new ListNode;
        

        auto ptr1 = list1;
        auto ptr2 = list2;
        if (ptr1 == NULL && ptr2 == NULL) return nullptr;

        auto curr = head;

        while(ptr1 != NULL && ptr2 != NULL) {
            
            if (ptr1->val >= ptr2->val) {
                curr->val = ptr2->val;
                ptr2 = ptr2->next;
            } else {
                curr->val = ptr1->val;
                ptr1 = ptr1->next;
            }
            
            curr->next = new ListNode;
            curr = curr->next;
        }

        if (ptr1 == NULL) {
            curr->val = ptr2->val;
            curr->next = ptr2->next;
        }
        if (ptr2 == NULL) {
            curr->val = ptr1->val;
            curr->next = ptr1->next;
        }
        return head;
    }
};
