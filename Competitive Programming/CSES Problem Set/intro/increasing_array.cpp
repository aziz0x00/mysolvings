#include <bits/stdc++.h>
#define ul unsigned long int

using namespace std;


int main()
{
  ios::sync_with_stdio(0);
  cin.tie(0);

  ul n, c = 0;
  cin >> n;

  ul a0, a;
  cin >> a0;
  while(cin >> a)
  {
    if (a0 > a)
      c += a0 - a;
    else
      a0 = a;
  }
  cout << c << "\n";
}

