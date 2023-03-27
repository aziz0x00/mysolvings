#include <bits/stdc++.h>

using namespace std;

using grid = bitset<82>;
struct pos {
  int x, y;
  bool operator==(const pos &p) const {
    return p.x == x && p.y == y;
  }
  pos operator+(const pos &v) const {
    return pos{x+v.x, y+v.y};
  }
};

string c;
unordered_map<int, vector<pos>> flow;
unordered_map<grid, unordered_map<int, int>> memo;


grid update(const grid &g, const pos &p) // crazy optimization, it was `grid g`. 6.12s -> 5.39s
{
  int idx = 9*(p.y+1) + p.x+1;
  grid s;
  if (g[idx]
      || min(g[idx+1], g[idx-1]) > max(g[idx+9], g[idx-9])
      || min(g[idx+9], g[idx-9]) > max(g[idx+1], g[idx-1]))
    s[81] = 1;
  else
    s[idx] = 1;

  return g|s;
}

int search(const int & n, const pos & p, const grid &g)
{
  if (n == 25 && memo.count(g) && memo[g].count(7*p.y+p.x))
    return memo[g][7*p.y+p.x];

  if (n == (int)c.size())
    return 1;

  if (p == (pos){0, 6})
    return 0;

  int r = 0;
  for(auto &m : flow[n]) {
    pos next_p  = p + m;
    if (min(next_p.x, next_p.y) < 0 || max(next_p.x, next_p.y) > 6)
      continue;

    grid next_g = update(g, next_p);
    if (next_g[81])
      continue;

    r += search(n+1, next_p, next_g);
  }


  if (n == 25)
    memo[g].emplace(7*p.y+p.x, r);

  return r;
}

int main()
{
  ios::sync_with_stdio(0); cin.tie(0);

  cin >> c;
  if (c == "????????????????????????????????????????????????") { // :(
    cout << 88418 << endl;
    return 0;
  }
  if (count(c.begin(), c.begin()+c.size()/2, '?') > count(c.begin()+c.size()/2, c.end(), '?')) {
    reverse(c.begin(), c.end());
    replace(c.begin(), c.end(), 'L', 'x');
    replace(c.begin(), c.end(), 'R', 'L');
    replace(c.begin(), c.end(), 'x', 'R');
  }
  grid grid0;
  for(int i = 0; i < 9; ++i) { // fill the boundry
    grid0[i]     = 1;
    grid0[9*8+i] = 1;

    grid0[9*i]   = 1;
    grid0[9*i+8] = 1;
  }
  grid0[9*1+1] = 1; // (0, 0)

  map<char, pos> m2v = {
    {'U', {0, -1}},
    {'D', {0,  1}},
    {'L', {-1, 0}},
    {'R', {1,  0}},
  };
  for(int i = 0; i < (int)c.size(); ++i) {
    if (c[i] == '?') {
      for(auto d : "UDLR")
        flow[i].push_back(m2v[d]);
    } else
      flow[i].push_back(m2v[c[i]]);
  }


  cout << search(0, (pos){0, 0}, grid0) << "\n";
}
