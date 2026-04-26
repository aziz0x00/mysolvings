class Solution {
public:
    vector<vector<int>> threeSum(vector<int>& nums) {
        vector<vector<int>> ans;
        const int n = nums.size();

        map<int, int> inv_nums;

        sort(nums.begin(), nums.end());

        for (int i = 0; i < n; ++i) {
            inv_nums[nums[i]] = i; // max preimage will be put here, that's why it works
        }

        set<pair<int, int>> elim;

        for (int i = 0; i < n; ++i) {
            for (int j = i + 1; j < n; ++j) {
                int c = -nums[i] - nums[j];
                if (inv_nums.count(c)) {
                    int k = inv_nums[c];
                    if (k == i || k == j)
                        continue;

                    if (nums[j] > c) continue; // make sure nums[i] <= nums[j] <= c
                    if (elim.count({nums[i], c})) continue;

                    elim.insert({nums[i], c});
                    ans.push_back({nums[i], nums[j], c});
                }
            }
        }

        return ans;
    }
};
