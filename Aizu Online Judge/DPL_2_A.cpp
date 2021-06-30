#include <bits/stdc++.h>
using namespace std;

const int INF = 1000100100;

template <class T>
bool chmin(T &a, const T &b)
{
    if (b < a)
    {
        a = b;
        return 1;
    }
    return 0;
}

int V, E;
int G[20][20];
int dp[50000][20];

int rec(int S, int v)
{
    if (S == 0)
    {
        if (v == 0)
            return 0;
        else
            return INF;
    }
    if ((S & (1 << v)) == 0)
    {
        return INF;
    }
    int &ret = dp[S][v];
    if (ret != 0)
        return ret;
    ret = INF;
    for (int u = 0; u < V; u++)
    {
        chmin(ret, rec(S ^ (1 << v), u) + G[u][v]);
    }
    return ret;
}

int main()
{
    int V, E;
    cin >> V >> E;
    for (int i = 0; i < 20; i++)
    {
        for (int j = 0; j < 20; j++)
        {
            G[i][j] = INF;
        }
    }
    for (int i = 0; i < E; i++)
    {
        int s, t, d;
        cin >> s >> t >> d;
        G[s][t] = d;
    }
    int ans = rec((1 << V) - 1, 0);
    if (ans != INF)
        cout << ans << endl;
    else
        cout << -1 << endl;
    return 0;
}
