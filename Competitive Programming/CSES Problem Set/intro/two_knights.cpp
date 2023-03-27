#include <bits/stdc++.h>

using namespace std;
typedef long long ll;


int main()
{
  ios::sync_with_stdio(0);
  cin.tie(0);

  int n;
  cin >> n;

  ll a = 0;
  for (ll i = 1; i <= n; i++)
  {
    if (i < 3) a = (i*i)*(i*i-1)>>1;
    else {
      a += (i-1)*(2*i-1) - 2;        // both knigths are on the new border
      a += (i-3)*(i-3)*(2*i-1);      // one on the border, and the other in (i-3) square
      a += (2*i-5)*(2*i-1 -2) -2 +2; // one on the border, and the other in (i-2) boundry
      a += (2*i-3)*(2*i-1 -2) +2;    // one on the border, and the other in (i-1) boundry
    }
    cout << a << "\n";
  }
}