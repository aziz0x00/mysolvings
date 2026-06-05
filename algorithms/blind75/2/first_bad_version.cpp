// The API isBadVersion is defined for you.
// bool isBadVersion(int version);

class Solution {
public:
    int firstBadVersion(int n) {
        int hi = n;
        int lo = 1;

        int mi;
        while (hi > lo) {
            mi = lo + (hi - lo) / 2;
            if (isBadVersion(mi)) {
                hi = mi;
            } else {
                lo = mi + 1;
            }
        }
        return hi;

        // int k = 0;
        // for(int b = n/2; b >= 1; b /= 2) {
        //     while (k <= n-b && !isBadVersion(b+k)) k += b;
        // }
        // return k+1;
    }
};
