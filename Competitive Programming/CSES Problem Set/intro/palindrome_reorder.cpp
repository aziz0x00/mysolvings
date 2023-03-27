#include <bits/stdc++.h>

using namespace std;
typedef long long ll;


int main()
{
  ios::sync_with_stdio(0);
  cin.tie(0);

  map<char, ll> ctr;

  string str, mid = "";
  cin >> str;
  for (auto a: str)
    ++ctr[a];
  str = "";
  for (auto a: ctr) {
    if (a.second % 2) {
      if (mid == "") {
        while(a.second--) mid += a.first;
        continue;
      }
      cout << "NO SOLUTION";
      return 0;
    }
    for (; a.second; a.second -= 2) str += a.first;
  }
  cout << str;
  reverse(str.begin(), str.end());
  if (mid != "") cout << mid;
  cout << str << "\n";
}