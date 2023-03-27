#include <bits/stdc++.h>
#define ul unsigned long int

using namespace std;


int main()
{
  ios::sync_with_stdio(0);
  cin.tie(0);

  ul n;
  cin >> n;

  if (n == 1)
  {
    cout << "1\n";
    return 0;
  }
  if (n <= 3)
  {
    cout << "NO SOLUTION\n";
    return 0;
  }

  for (ul i = 1; i < n/2+1; ++i)
    cout << 2*i << " ";
  for (ul i = 1; i < (n+1)/2+1; ++i)
    cout << 2*i-1 << " ";
}