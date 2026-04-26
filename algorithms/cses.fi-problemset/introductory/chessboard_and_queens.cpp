#include <bits/stdc++.h>

using namespace std;
typedef long long ll;


struct state {
  bool diag[16], gaid[16], col[8]; // gaid as of "diag" reversed
};

ll search(int n, state s, bool board[8][8])
{
  auto line = board[n];

  ll r = 0;
  for(int i = 0; i < 8; ++i)
  {
    if (line[i] || s.col[i] || s.diag[7-n+i] || s.gaid[n+i])
      continue;

    if (n == 0)
      return 1;
    else {
      state next_s       = s;
      next_s.col[i]      = 1;
      next_s.diag[7-n+i] = 1;
      next_s.gaid[n+i]   = 1;
      r += search(n-1, next_s, board);
    }
  }
  return r;
}

int main()
{
  ios::sync_with_stdio(0);
  cin.tie(0); cout.tie(0);

  bool board[8][8];
  string l;
  for(int i = 0; i < 8; ++i) {
    cin >> l;
    for(int j = 0; j < 8; ++j) board[i][j] = l[j] == '*';
  }

  cout << search(7, (state){}, board) << "\n";
}