#include <bits/stdc++.h>

using namespace std;


int main()
{
  ios::sync_with_stdio(0);
  cin.tie(0);

  long long int n, r, j;
  cin >> n;
  r = n*(n+1) >> 1;

  for (int i = 1; i < n; i++) {
    cin >> j;
    r -= j;
  }

  cout << r << "\n";
}