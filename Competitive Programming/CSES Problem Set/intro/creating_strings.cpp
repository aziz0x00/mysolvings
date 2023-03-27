#include <bits/stdc++.h>

using namespace std;
typedef long long ll;

void spam(string head, map<char, int> ctr, int n)
{
  if ((int)head.size() == n) {
    cout << head << "\n";
    return;
  }

  for(auto a : ctr)
  {
    if (a.second == 0)
      continue;

    --ctr[a.first];
    spam(head + a.first, ctr, n);
    ++ctr[a.first];
  }
}


int main()
{
  ios::sync_with_stdio(0);
  cin.tie(0);

  string str;
  cin >> str;

  int n = str.size();
  map<char, int> ctr;
  for(auto a : str) ++ctr[a];

  // counting
  int c = 1;
  for(int i = n; i > 0; --i) c *= i;
  for(auto a : ctr)
    for(int i = a.second; i > 0; --i) c /= i;
  cout << c << "\n";

  // spamming all of them
  spam("", ctr, n);
}