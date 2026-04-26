# The isBadVersion API is already defined for you.
# def isBadVersion(version: int) -> bool:

class Solution:
    def firstBadVersion(self, n: int) -> int:
        
        # hi = n
        # lo = 0
        # while hi > lo:
        #     mi = lo + (hi - lo) // 2
        #     if isBadVersion(mi):
        #         hi = mi
        #     else:
        #         lo = mi + 1
        # return hi

        b = n//2
        k = 0
        while b >= 1:
            while k+b <= n and not isBadVersion(b+k):
                k += b
            b //= 2
        
        return k+1
