// Kadane
class Solution {
public:
    int maxSubArray(vector<int>& nums) {
        int best = nums[0];
        int sum = 0;
        for(auto &num : nums) {
            sum = max(num, num+sum);
            best = max(sum, best);
        }
        return best;
    }
};
