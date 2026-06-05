class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        m = {}
        for i, x in enumerate(nums):
            complement = target - x
            if complement in m:
                return [i, m[complement]]
            m[x] = i
