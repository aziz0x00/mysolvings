// clean solution

class Solution {
public:
    vector<vector<int>> insert(vector<vector<int>>& intervals, vector<int>& newInterval) {
        vector<vector<int>> ans;

        int s = newInterval[0], e = newInterval[1];

        int i = 0, n = size(intervals);

        while (i < n && intervals[i][1] < s)
            ans.push_back(intervals[i++]);

        while (i < n && e >= intervals[i][0]) {
            s = min(s, intervals[i][0]);
            e = max(e, intervals[i][1]);
            ++i;
        }

        ans.push_back({s, e});

        while (i < n)
            ans.push_back(intervals[i++]);
        
        return ans;
    }
};
