#include <bits/stdc++.h>
using namespace std;

vector<vector<int> > nbrs(100);
vector<int> ds(100), fs(100);

int dfs(int u, int t)
{
    ds[u] = ++t;
    int l = nbrs[u].size();
    for (int i = 0; i < l; i++)
    {
        int &v = nbrs[u][i];
        if (!ds[v])
            t = dfs(v, t);
    }
    return (fs[u] = ++t);
}

int main()
{
    int n;
    cin >> n;

    for (int i = 0; i < n; i++)
    {
        int u, k;
        cin >> u >> k;
        u--;
        for (int j = 0; j < k; j++)
        {
            int a;
            cin >> a;
            nbrs[u].push_back(a);
        }
    }

    int t = 0;
    for (int i = 0; i < n; i++)
    {
        if (!ds[i])
            t = dfs(i, t);
    }

    for (int i = 0; i < n; i++)
    {
        cout << i << " " << ds[i] << " " << fs[i] << endl;
    }
}
