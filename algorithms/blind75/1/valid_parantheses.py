class Solution:
    def isValid(self, s: str) -> bool:
        closing = {k: v for v, k in zip("{[(", "}])")}
        stack = []
        for c in s:
            if c in closing:
                if len(stack) == 0:
                    return False
                c = closing[c]
                if c != stack.pop():
                    return False
            else:
                stack.append(c)
        
        return len(stack) == 0
        
