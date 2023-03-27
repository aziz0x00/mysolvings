#include <bits/stdc++.h>

using namespace std;
typedef unsigned long long ll;


int main()
{
  ios::sync_with_stdio(0);
  cin.tie(0);

  ll n, s = 1;
  cin >> n;
  while (n--)
    s = (s << 1) % 1000000007;
  cout << s << "\n";
}
