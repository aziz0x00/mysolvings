# python lists slower than just arrays ?
class Solution:
    def canConstruct(self, ransomNote: str, magazine: str) -> bool:
        a, b = Counter(ransomNote), Counter(magazine)
        return a & b == a
