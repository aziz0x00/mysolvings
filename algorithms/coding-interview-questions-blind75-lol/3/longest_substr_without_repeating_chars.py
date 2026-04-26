class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        i = 0
        chars = [-1]*0xFF
        ans = 0
        for j in range(len(s)):
            if chars[ord(s[j])] >= i:
                i = chars[ord(s[j])] + 1
            chars[ord(s[j])] = j
            ans = max(ans, j - i + 1)
        
        return ans
