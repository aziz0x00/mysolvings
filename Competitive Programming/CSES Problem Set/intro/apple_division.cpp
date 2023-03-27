#include <bits/stdc++.h>

using namespace std;
typedef long long ll;

int main()
{
  ios::sync_with_stdio(0);
  cin.tie(0);

  int n;
  cin >> n;

  ll p[n];
  for(int i = 0; i < n; ++i) cin >> p[i];

  ll mi = numeric_limits<ll>::max(), s;
  if (n == 1) mi = p[0];
  for (int i = 1; i < (1 << n)-1; ++i)
  {
    s = 0;
    for(int idx = 0; idx < n; ++idx)
      s += pow(-1, 1&(i >> idx))*p[idx];

    mi = min(mi, abs(s));
    if (mi == 0) break;
  }
  cout << mi << "\n";
}