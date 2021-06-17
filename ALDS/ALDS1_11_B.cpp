#include <bits/stdc++.h>
using namespace std;

int timestamp = 0;

void traverse(vector<vector<int> > &v, vector<vector<int> > &ans, int t)
{
    timestamp++;
    ans[t][0] = timestamp;
    for (int i = v[t].size() - 1; i >= 0; i--)
    {
        traverse(v, ans, i);
    }
    timestamp++;
    ans[t][1] = timestamp;
    return;
}

int main()
{
    int n;
    cin >> n;
    vector<vector<int> > v(n);

    for (int i = 0; i < n; i++)
    {
        int u, k;
        cin >> u >> k;
        for (int j = 0; j < k; j++)
        {
            int a;
            cin >> a;
            v[u].push_back(a);
        }
    }

    vector<vector<int> > ans(n, vector<int>(2));

    traverse(v, ans, 0);

    for (int i = 0; i < n; i++)
    {
        cout << v[i][0] << " " << v[i][1] << endl;
    }
}
