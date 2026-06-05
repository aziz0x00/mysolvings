#include <bits/stdc++.h>

using namespace std;
typedef long long ll;


int main()
{
  ios::sync_with_stdio(0);
  cin.tie(0);

  ll T, a, b, ma, mi, d;
  cin >> T;

  while (T--)
  {
    cin >> a >> b;
    ma = max(a, b);
    mi = min(a, b);

    d = ma - mi;

    if (mi-d >= 0 && (mi-d)%3 == 0)
      cout << "YES\n";
    else
      cout << "NO\n";
  }
}