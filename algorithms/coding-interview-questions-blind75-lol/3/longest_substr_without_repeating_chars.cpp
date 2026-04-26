// first try
class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        int n = s.size();

        int i = 0, j = 0;
        set<char> ctr;

        int ans = 0;

        while (j < n && i < n) {
            printf("%d - %d\n", i, j);
            if (ctr.count(s[j])) {
                ans = max(ans, j - i);
                while (s[i] != s[j]) {
                    ctr.erase(s[i]);
                    ++i;
                }
                ++i;
            } else {
                ctr.insert(s[j]);
            }
            ++j;
        }
        ans = max(ans, j - i);

        return ans;
    }
};

// second attempt
class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        int n = s.size();

        int i = 0, j = 0;
        // set<char> cSet;
        bool chars[0xFF];
        memset(chars, 0, size(chars));

        int ans = 0;

        for(; j < n; ++j) {
            if (chars[s[j]]) {
                while (s[i] != s[j]) {
                    chars[s[i]] = 0;
                    ++i;
                }
                ++i;
            } else {
                chars[s[j]] = 1;
                ans = max(ans, j - i + 1);
            }
        }

        return ans;
    }
};

// best sol on a lc post (YO MY LAST SOL IS FASTER)
class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        int n = s.size();

        vector<int> chars(0xff, -1);

        int ans = 0;

        for(int i = 0, j = 0; j < n; ++j) {
            if (chars[s[j]] >= i) {
                i = chars[s[j]] + 1;
            }
            chars[s[j]] = j;
            ans = max(ans, j - i + 1);
        }

        return ans;
    }
};
