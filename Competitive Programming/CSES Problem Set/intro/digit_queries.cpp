#include <bits/stdc++.h>

using namespace std;
typedef unsigned long long ull;


ull s(int n)
{
  // we need to calculate 9*(sum[1,n] k*10^(k-1)), so:
  // sum k*x^(k-1) = d( sum x^k ) = d( (x^(n+1)-1)/(x-1) ) = (n+1)x^n/(x-1) - (x^(n+1)-1)/(x-1)^2
  return (n+1)*((ull)pow(10, n)) - ((ull)pow(10, n+1)-1)/9;
}

int main()
{
  ios::sync_with_stdio(0);
  cin.tie(0);

  int T;
  cin >> T;

  ull k;
  while(cin >> k)
  {
    int n = 1;
    while(s(n) < k) ++n; // feasible, because log10(k) < 19

    ull m = (ull)pow(10, n-1) + (k-s(n-1)-1)/n; // the number itself
    int i = (k-s(n-1)-1) % n; // the index of the digit
    cout << (m / ((ull)pow(10,n-i-1))) % 10 << "\n";
  }
}