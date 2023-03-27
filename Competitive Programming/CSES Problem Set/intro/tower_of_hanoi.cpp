#include <bits/stdc++.h>

using namespace std;
typedef long long ll;


void tower(int n, int start, int end)
{
  if (n == 0) return;

  int free = 3*2 - start - end;
  tower(n-1, start, free);
  printf("%d %d\n", start, end);
  tower(n-1, free, end);
}

int main()
{
  ios::sync_with_stdio(0);
  cin.tie(0); cout.tie(0);

  int n;
  cin >> n;

  printf("%d\n", (1 << n) -1);

  tower(n, 1, 3);
}