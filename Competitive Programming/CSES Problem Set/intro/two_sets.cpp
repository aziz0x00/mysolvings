#include <bits/stdc++.h>

using namespace std;
typedef long long ll;


int main()
{
  ios::sync_with_stdio(0);
  cin.tie(0); cout.tie(0);

  ll n;
  cin >> n;

  if ((n*(n+1)>>1) & 1) {
    cout << "NO\n";
    return 0;
  }
  cout << "YES\n";
  ll s = n*(n+1) >> 2; // the sum of each subset
  ll m = sqrt(2*s); // 1+2+...+m =< s < 1+...+m+(m+1)
  if ((m*(m+1)>>1) > s) --m; // ?????WHYYY

  cout << m << "\n";
  ll a = 0;
  for (ll i = 1; i <= m; i++)
  {
    if (s == (m*(m+1)>>1)-i+m+1) {
      a = i;
      cout << m+1 << " ";
    } else
      cout << i << " ";
  }
  cout << "\n" << n-m << "\n";
  cout << (a ? a : m+1);
  for (ll i = m+2; i <= n; i++)
    cout << " " << i;
}