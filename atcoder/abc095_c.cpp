#include <bits/stdc++.h>
using namespace std;

int main()
{
    int A, B, C, X, Y;
    cin >> A >> B >> C >> X >> Y;
    int ans = 100100100100;

    for (int i = 0; i <= 100000; i++)
    {
        ans = min(ans, (i * 2 * C) + (max(0, X - i) * A) + (max(0, Y - i) * B));
    }
    cout << ans << endl;
}