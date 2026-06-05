# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def mergeTwoLists(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
        if list1 == list2 == None:
            return None
        
        head = ListNode()

        ptr1 = list1
        ptr2 = list2
        curr = head
        while ptr1 != None and ptr2 != None:
            if ptr1.val >= ptr2.val:
                curr.next = ptr2
                ptr2 = ptr2.next
            else:
                curr.next = ptr1
                ptr1 = ptr1.next
            
            curr = curr.next

        tail = ptr1 or ptr2
        if tail:
            curr.next = tail
        
        return head.next
        
