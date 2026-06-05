// it was Fibonacci and missed it
class Solution {
public:
    int climbStairs(int n) {
        int ans = 1; // all 1's case

        for(int i = 1; i <= n/2; ++i) {
            unsigned long long k = 1;
            int l = max(n-2*i, i);
            for(int j = l+1, mi = 1, mv = 1; j <= n-i; ++j) { // calculate (n - i)!/(i!(n-2*i)!)
                if (mv == 1 && mi < n-i-l) mv = ++mi;
                k *= j;
                if (mv != 1) {
                    int g = gcd(k, mv);
                    mv /= g;
                    k /= g;
                }
            }
            // for(int j = l+1; j <= n-i; ++j) k *= j;
            // for(int j = 1; j <= n-i-l; ++j) k /= j;
            ans += k;
        }
        
        return ans;
    }
};

/*
    n = 2*i + m, i = 0..(n/2)
    for each i, a climbStair is made by successive choices from i 2's and m 1's which count in total as (i+m)!/(i!m!)
*/
