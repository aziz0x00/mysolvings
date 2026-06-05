class Solution:
    def search(self, nums: List[int], target: int) -> int:
        
        lo, hi = 0, len(nums)-1

        while hi >= lo:
            mi = (hi + lo) // 2
            if nums[mi] > target:
                hi = mi - 1
            elif nums[mi] < target:
                lo = mi + 1
            else:
                return mi
        
        return -1
