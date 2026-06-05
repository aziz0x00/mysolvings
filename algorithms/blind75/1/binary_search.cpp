class Solution {
public:
    int search(vector<int>& nums, int target) {
        int hi = size(nums)-1;
        int lo = 0;
        int mi = 0;
        while (hi >= lo) {
            mi = (hi + lo) / 2;
            if (nums[mi] == target) {
                return mi;
            } else if (nums[mi] > target) {
                hi = mi - 1;
            } else {
                lo = mi + 1;
            }
        }
        

        return -1;
    }
};
