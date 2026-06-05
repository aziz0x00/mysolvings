#include <bits/stdc++.h>

using namespace std;


int main()
{
  ios::sync_with_stdio(0);
  cin.tie(0);

  int n;
  cin >> n;

  set<int> nset;
  int a;
  while (cin >> a) nset.insert(a);

  cout << nset.size() << "\n";
}