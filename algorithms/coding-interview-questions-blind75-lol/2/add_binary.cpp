// i did the trivial approach, but this is better
class Solution {
public:
    string addBinary(string a, string b) {
        string ans = "";

        int carry = 0;

        int i = a.length() - 1, j = b.length() - 1;

        while(i >= 0 || j >= 0 || carry) {
            if (i >= 0)
                carry += a[i--] - '0';
            if (j >= 0)
                carry += b[j--] - '0';
            ans = (char)(carry%2 + '0') + ans;
            carry /= 2;
        }
        return ans;
    }
};
