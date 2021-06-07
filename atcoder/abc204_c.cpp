#include <bits/stdc++.h>
using namespace std;
 
int main() {
  int N, M;
  cin >> N >> M;
  vector<vector<int>> G(N);
  
  for (int i = 0; i < M; i++) {
    int a, b;
    cin >> a >> b;
    a--;
    b--;
    G[a].push_back(b);
  }
  
  int ans = 0;
  for (int i = 0; i < N; i++) {
    vector<int> used(N);
    queue<int> q;
    q.push(i);
    ans++;
    used[i] = 1;
    while (q.size()) {
      int t = q.front(); q.pop();
      for (auto u: G[t]) {
	    if (used[u] == 1) continue;
        ans++;
        used[u] = 1;
        q.push(u);
      }
    }
  }
 
  
  cout << ans << endl;
  
  return 0;
}  