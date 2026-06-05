class Solution {
public:
    int maxProfit(vector<int>& prices) {

        const int N = size(prices);

        int max_profit = 0;
        int mi = prices[0];

        for (int i = 1; i < N; ++i) {
            max_profit = max(max_profit, prices[i] - mi);
            mi = min(mi, prices[i]);
        }

        return max_profit;
    }
};
