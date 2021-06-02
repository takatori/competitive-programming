#include <bits/stdc++.h>
using namespace std;

int main()
{
    int N, M;
    cin >> N >> M;
    vector<vector<int> > S(M, vector<int>(N));
    vector<int> P(M);

    for (int i = 0; i < M; i++)
    {
        int k;
        cin >> k;
        for (int j = 0; j < k; j++)
        {
            int a;
            cin >> a;
            S[i][a - 1] = 1;
        }
    }
    for (int i = 0; i < M; i++)
        cin >> P[i];

    int max = (2 << (N - 1));
    int ans = 0;

    for (int i = 0; i < max; i++)
    {
        bool ok = true;
        for (int j = 0; j < M; j++)
        {
            int count = 0;
            for (int k = 0; k < N; k++)
            {
                if (S[j][k] == 1 && (i >> k & 1) == 1)
                    count++;
            }
            if (count % 2 != P[j])
                ok = false;
        }
        if (ok)
            ans++;
    }

    cout << ans << endl;
}