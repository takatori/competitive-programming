#include <bits/stdc++.h>
using namespace std;

int w[300];

int rec(int l, int r, vector<vector<int> > &dp)
{
    if (dp[l][r] != -1)
        return dp[l][r];
    if (abs(l - r) <= 1)
        return 0;
    int res = 0;
    if (abs(w[l] - w[r - 1]) <= 1 && rec(l + 1, r - 1, dp) == r - l - 2)
    {
        res = r - l;
    }
    for (int mid = l + 1; mid <= r - 1; mid++)
    {
        res = max(res, rec(l, mid, dp) + rec(mid, r, dp));
    }
    return dp[l][r] = res;
}

int main()
{
    int n;
    cin >> n;
    while (n > 0)
    {
        for (int i = 0; i < n; i++)
            cin >> w[i];
        vector<vector<int> > dp(n + 2, vector<int>(n + 2, -1));
        cout << rec(0, n, dp) << endl;
        cin >> n;
    }
}