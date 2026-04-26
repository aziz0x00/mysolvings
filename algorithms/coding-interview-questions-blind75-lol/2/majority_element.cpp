class Solution {
public:
    int majorityElement(vector<int>& nums) {
        int m = size(nums) / 2;
        unordered_map<int, int> mp;

        for(auto& num : nums) {
            if (!mp.count(num)) {
                mp[num] = 1;
            } else {
                ++mp[num];
            }
            if (mp[num] > m) return num;
        }
        return 0;
    }
};


/* Boyer-Moore voting algorithm

class Solution {
public:
    int majorityElement(vector<int>& nums) {
        // int m = size(nums) / 2;
        // unordered_map<int, int> mp;

        int count = 0;
        int ans;

        for(auto& num : nums) {
            if (count == 0) ans = num; // this only happens when left #{.!=num} and #(.==num) are equal. and then we choose something not 
            if (num == ans) ++count;
            else --count;
        }

        return ans;
    }
};

*/
