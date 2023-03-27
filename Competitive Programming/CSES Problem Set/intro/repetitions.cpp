#include <bits/stdc++.h>

using namespace std;


int main()
{
  int n = 1, l;
  string str;
  cin >> str;

  for (int i = 1; i < str.size(); i++)
  {
    l = i-1;
    while (str[l] == str[i]) i++;
    n = max(n, i-l);
  }

  cout << n << "\n";
}