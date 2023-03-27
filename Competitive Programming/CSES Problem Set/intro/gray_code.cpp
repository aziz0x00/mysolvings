#include <bits/stdc++.h>

using namespace std;
typedef long long ll;

void gcode(string head, unsigned int n, bool s = 0)
{
  if (n == 0) {
    cout << head << "\n";
    return;
  }
  gcode(head + to_string(s),  n-1, 0);
  gcode(head + to_string(!s), n-1, 1);
}

int main()
{
  ios::sync_with_stdio(0);
  cin.tie(0);

  unsigned int n;
  cin >> n;

  gcode("", n);
}