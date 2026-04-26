#include <bits/stdc++.h>

using namespace std;
typedef long long ll;

int main() {
    ios::sync_with_stdio(0);
    cin.tie(0);

    ll n, x;

    cin >> n >> x;
    ll p[n];
    for (int i = 0; i < n; ++i) cin >> p[i];
    sort(p, p + n);

    ll ans = 0;
    ll i = 0;
    ll j = n-1;
    while ( i < j) {
        ++ans;
        if (p[i] + p[j--] <= x) { // the strategy is to pick from lighest and heaviest.
            ++i;
        }
    }
    if (i == j) ++ans;

    cout << ans << endl;
}
