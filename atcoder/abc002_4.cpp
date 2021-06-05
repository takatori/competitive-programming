#include <bits/stdc++.h>
using namespace std;
 
int main() {
  int n, m;
  cin >> n >> m;
  vector<bitset<12>> x(n);
  for (int i = 0; i < n; i++) 
    	x[i][i] = true;
  for (int i = 0; i < m; i++) {
    int a, b;
    cin >> a >> b;
    x[a-1][b-1] = x[b-1][a-1] = true;
  }
  
  int ret = 0;
  for (int i = 0; i < (1<<n); i++) {
    bitset<12> bs(i);
    bool ok = true;
    for (int j = 0; j < n; j++) {
      if (bs[j] && (bs & ~x[j]).any()) ok = false;
    }
    if (ok) 
      ret = max(ret, (int)bs.count());
  }
  cout << ret << endl;
  return 0;
}