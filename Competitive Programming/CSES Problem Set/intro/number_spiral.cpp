#include <bits/stdc++.h>

using namespace std;


int main()
{
  ios::sync_with_stdio(0);
  cin.tie(0);
  cout.tie(0);

  long long T, x, y, mi, ma;
  cin >> T;

  while (T--)
  {
    cin >> y >> x;
    mi = min(x, y);
    ma = max(x, y);

    // if ( the boundry is odd and we are searching on the vertical )
    // OR ( the boundry is even and we are searching on the horizontal )
    if ((ma & 1 && y < x) || !(ma & 1 || y < x))
      cout << ma*ma - mi + 1 << "\n";
    else
      cout << (ma-1)*(ma-1) + mi << "\n";
  }
}