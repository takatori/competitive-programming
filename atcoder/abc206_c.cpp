#include <bits/stdc++.h>
using namespace std;

int main() {

  int N;
  cin >> N;
  vector<int> A(N);
  long long ans = 0;	
  for (int i = 0; i < N; i++) cin >> a[i];

  unordered_map<int, int> cnt;
  for (int i = 0; i < N; i++) {
    ans += i - cnt[A[i]];
    cnt[A[i]]++;
  }
  
  cout << ans << endl;
   
}