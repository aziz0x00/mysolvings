/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution { // one turtle was enough
public:
    bool hasCycle(ListNode *head) {
        ListNode* turtle1 = head;
        ListNode* turtle2 = head;
        ListNode* rabbit = head;

        if (head == NULL || head->next == NULL) return false;

        for (;;) {
            if (rabbit->next == NULL || rabbit->next->next == NULL) return false;
            rabbit = rabbit->next->next;

            if (rabbit == turtle1 || rabbit == turtle2) return true;

            turtle1 = turtle1->next;
            turtle2 = turtle2->next;
        }
    }
};
