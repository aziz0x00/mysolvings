#include <bits/stdc++.h>

using namespace std;
typedef unsigned long long ll;


int main()
{
  ios::sync_with_stdio(0);
  cin.tie(0);

  ll n;
  cin >> n;

  ll a = 0;
  while (n /= 5)
    a += n;

  cout << a << "\n";
}