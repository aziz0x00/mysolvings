#include <bits/stdc++.h>

using namespace std;


int main()
{
  long long int n;
  cin >> n;
  for (; n != 1; n = n&1 ? 3*n+1 : n>>1)
    cout << n << " ";

  cout << "1 \n";
}