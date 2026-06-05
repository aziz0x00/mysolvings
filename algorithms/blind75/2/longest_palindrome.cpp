class Solution {
public:
    int longestPalindrome(string s) {
        int chars[26 + 26];
        memset(chars, 0, size(chars));
        for(auto&c : s) {
            if (c >= 'a') {
                chars[c - 'a']++;
            } else {
                chars[26 + c - 'A']++;
            }
        }
        bool odd = 0;
        int accum = 0;
        for(int i = 0; i < size(chars); ++i) {
            if (chars[i] % 2) {
                accum += chars[i] - 1;
                odd = 1;
            } else {
                accum += chars[i];
            }
        }
        return accum + odd;
    }
};
/* // better:
class Solution {
public:
    int longestPalindrome(string s) {
        int chars[26 + 26];
        memset(chars, 0, size(chars));
        int oddCount = 0;
        for(auto&c : s) {
            int pos;
            if (c >= 'a') {
                pos = c - 'a';
            } else {
                pos = 26 + c - 'A';
            }
            chars[pos]++;
            if (chars[pos] % 2) {
                oddCount++;
            } else {
                oddCount--;
            }
        }
        if (oddCount > 1) return size(s) - oddCount + 1;
        return size(s);
    }
};
*/
