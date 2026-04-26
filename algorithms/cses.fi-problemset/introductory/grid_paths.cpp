#include <bits/stdc++.h>

using namespace std;

bool g[9 * 9];
vector<int> flow[9 * 9];

int search(const int p, const int n) {

  if (p == 9 * 7 + 1)
    return n == 48;

  int r = 0;
  g[p] = true;
  for (auto &m : flow[n]) {
    int q = p + m;

    // check already visited or out-of-bound
    if (g[q])
      continue;

    bool a = (g[q - 1] & g[q + 1]) > (g[q - 9] | g[q + 9]);
    bool b = (g[q - 1] | g[q + 1]) < (g[q - 9] & g[q + 9]);
    // check dead-end
    if (a ^ b)
      continue;

    r += search(q, n + 1);
  }
  g[p] = false;

  return r;
}

int main() {
  ios::sync_with_stdio(0);
  cin.tie(0);

  string path;
  cin >> path;
  // if (count(path.begin(), path.begin()+path.size()/2, '?') >
  // count(path.begin()+path.size()/2, path.end(), '?')) {
  //     reverse(path.begin(), path.end());
  //     replace(path.begin(), path.end(), 'L', 'x');
  //     replace(path.begin(), path.end(), 'R', 'L');
  //     replace(path.begin(), path.end(), 'x', 'R');
  // }
  map<char, int> m2v = {{'U', 0}, {'D', 1}, {'L', 2}, {'R', 3}};
  const int dz[] = {-9, 9, -1, 1};

  for (int i = 0; i < 49; ++i)
    if (path[i] == '?')
      flow[i] = {-9, 9, -1, 1};
    else
      flow[i] = {dz[m2v[path[i]]]};

  for (int i = 0; i < 9; ++i)
    for (int j = 0; j < 9; ++j)
      g[i * 9 + j] = (min(i, j) == 0 or max(i, j) == 8);

  cout << search(1 + 1 * 9, 0) << "\n";
}
