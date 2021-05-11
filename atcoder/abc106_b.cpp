#include <bits/stdc++.h>
using namespace std;
 
 
int divisor(long long n) {
  vector<long long> ret;
  for (long long i = 1; i * i <= n; i++) {
    if (n % i == 0) {
      ret.push_back(i);
      if (i * i != n) ret.push_back(n/i);
    }
  }
  return ret.size();
}
 
 
int main() {
  long long N;
  cin >> N;
 
  int ans = 0;
  for (int i = 105; i <= N; i += 2) {
    if (divisor(i) == 8) ans++;
  }
 
  cout << ans << endl;
}