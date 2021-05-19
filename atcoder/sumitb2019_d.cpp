#include <bits/stdc++.h>
using namespace std;

int main()
{
    int N;
    string S;
    cin >> N >> S;

    int ans = 0;

    for (int i = 0; i < 1000; i++)
    {
        int a = i % 10;
        int b = (i / 10) % 10;
        int c = (i / 100) % 10;
        auto c_pos = find(S.begin(), S.end(), '0' + c);
        auto b_pos = find(c_pos + 1, S.end(), '0' + b);
        auto a_pos = find(b_pos + 1, S.end(), '0' + a);
        if (a_pos != S.end() && b_pos != S.end() && c_pos != S.end())
            ans++;
    }

    cout << ans << endl;
}