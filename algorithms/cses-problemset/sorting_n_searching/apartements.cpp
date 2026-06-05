#include <bits/stdc++.h>

using namespace std;
typedef long long ll;

int main() {
    ios::sync_with_stdio(0);
    cin.tie(0);

    int ans = 0;
    int n, m;
    ll k;
    cin >> n >> m >> k;

    ll appli[n], apart[m];
    // sort them to go through them linearly for comparaison
    for (int i = 0; i < n; ++i) cin >> appli[i];
    sort(appli, appli + n);
    for (int i = 0; i < m; ++i) cin >> apart[i];
    sort(apart, apart + m);

    int i = 0;
    for (auto ap : appli) {
        while (i < m && ap > apart[i]+k ) ++i; // skip too smaller apartements
        if (i == m) break;

        if (ap >= apart[i]-k) { // here is only reached when`ap <= apart[i]+k`,
            ++ans;              // so the other bound is what's left to be checked.

            // and only go to next apartement when the current is picked
            ++i;
        }
    }

    cout << ans << "\n";
}
