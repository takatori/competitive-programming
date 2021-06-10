#include <bits/stdc++.h>
using namespace std;

int main()
{
    int d, n, m;
    cin >> d;
    cin >> n;
    cin >> m;

    vector<int> N(n);
    for (int i = 0; i < n; i++)
        cin >> N[i];
    sort(N.begin(), N.end());

    int ans = 0;

    for (int i = 0; i < m; i++)
    {
        int M;
        cin >> M;

        auto iter = lower_bound(N.begin(), N.end(), M);
        int dist;
        if (iter != N.end())
        {
            dist = *iter - M;
        }
        else
        {
            dist = d - M;
        }

        ans += dist;
    }
    cout << ans << endl;
}