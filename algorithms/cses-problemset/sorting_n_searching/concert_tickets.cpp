#include <bits/stdc++.h>

using namespace std;

typedef long long ll;


int main() {
    int n, m;
    cin >> n >> m;

    ll h[n], t;

    for (int i = 0; i < n; ++i) cin >> h[i];
    sort(h, h+n);
    for (int i = 0; i < n; ++i) cout << h[i] << " ";
    cout << endl;


    bool picked[n];
    memset(picked, 0, n);
    int hi, lo, mi;
    while (cin >> t) {
        // bisect for least upper bound
        hi = n-1;
        lo = 0;
        while (hi - lo > 1) {
            mi = (hi + lo)/2;
            if (h[mi] <= t) {
                lo = mi;
            } else {
                hi = mi;
            }
        }
        mi = hi;
        printf("mi = %d, h[mi] = %lld\n", mi, h[mi]);
        while (picked[mi] && mi >= 0) mi--;

        if (mi == -1) {
            cout << -1 << endl;
        } else {
            cout << h[mi] << endl;
            picked[mi] = 1;
        }
    }
}
