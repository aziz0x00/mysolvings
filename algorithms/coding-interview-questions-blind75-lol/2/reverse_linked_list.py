# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if head == None: return
        ans = None
        while head:
            temp = head.next
            head.next = ans
            ans = head
            head = temp
        
        return ans

"""
# Solution that doesn't mutate head

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        ans = None
        while head:
            temp = ListNode()
            temp.next = ans
            ans = temp
            ans.val = head.val
            head = head.next
        
        return ans
"""
