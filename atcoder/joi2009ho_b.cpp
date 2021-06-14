#include <bits/stdc++.h>
using namespace std;

int main()
{
    int d, n, m;
    cin >> d;
    cin >> n;
    cin >> m;

    vector<int> in(n - 1);
    for (int i = 0; i < n - 1; i++)
        cin >> in[i];
    in.push_back(0);
    sort(in.begin(), in.end());

    int ans = 0;

    while (m--)
    {
        int a;
        cin >> a;

        int idx = lower_bound(in.begin(), in.end(), a) - in.begin();

        int s = 1 << 30;
        if (idx != n)
            s = in[idx] - a;
        else
            s = d - a;

        if (idx)
            s = min(s, a - in[idx - 1]);
        ans += s;
    }

    cout << ans << endl;
}