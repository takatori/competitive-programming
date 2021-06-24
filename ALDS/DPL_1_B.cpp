#include <bits/stdc++.h>
using namespace std;

int dp[110][10010];

int main()
{
    int N, W;
    cin >> N >> W;
    vector<int> v(N);
    vector<int> weight(N);

    for (int i = 0; i < N; i++)
        cin >> v[i] >> weight[i];
    for (int w = 0; w <= W; w++)
        dp[0][w] = 0;

    for (int i = 0; i < N; i++)
    {
        for (int w = 0; w <= W; w++)
        {
            if (w >= weight[i])
                dp[i + 1][w] = max(dp[i][w - weight[i]] + v[i], dp[i][w]);
            else
                dp[i + 1][w] = dp[i][w];
        }
    }

    cout << dp[N][W] << endl;
}