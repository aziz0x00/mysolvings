#include <bits/stdc++.h>

using namespace std;

class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        map<int, int> m;
        int complement;
        for(int i = 0; i < size(nums); ++i) {
            complement = target - nums[i];
            if (m.count(complement))
                return {i, m[complement]};
            m[nums[i]] = i;
        }
        return {};
    }
};
