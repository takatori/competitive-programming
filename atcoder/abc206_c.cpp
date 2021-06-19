#include <bits/stdc++.h>
using namespace std;

int main() {

  int N;
  cin >> N;
  unordered_map<int, long long> a;
  long long ans = N * (N-1) / 2;	

  for (int i = 0; i < N; i++) {
    int x;
    cin >> x;
    a[x]++;
  }
  
  for (auto t : a) {
    // cout << t.first << " " << t.second << endl;
    ans = ans - (t.second * (t.second-1) / 2);
  }
  
  cout << ans << endl;
   
}