#include <bits/stdc++.h>
using namespace std;

vector<vector<int> > nbrs(100);
vector<int> bs(101, -1);
queue<int> q;

int main()
{
    int n;
    cin >> n;

    for (int i = 0; i < n; i++)
    {
        int u, k;
        cin >> u >> k;
        u;
        for (int j = 0; j < k; j++)
        {
            int a;
            cin >> a;
            nbrs[u].push_back(a);
        }
    }

    q.push(1);
    bs[1] = 0;
    while (!q.empty())
    {
        int current = q.front();
        q.pop();
        for (int vertex : nbrs[current])
        {
            if (bs[vertex] != -1)
                continue;
            bs[vertex] = bs[current] + 1;
            q.push(vertex);
        }
    }

    for (int i = 1; i <= n; i++)
    {
        cout << i << " " << bs[i] << endl;
    }
}
