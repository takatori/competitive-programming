#include <bits/stdc++.h>
using namespace std;

string S;
int N;
string T = "ACGT";

int main()
{
    cin >> S;
    N = S.length();
    int ans = 0;
    for (int i = 0; i < N; i++)
    {
        for (int j = i; j < N; j++)
        {
            int ok = 1;
            for (int x = i; x < j + 1; x++)
            {
                if (T.find(S[x]) == string::npos)
                    ok = 0;
            }
            if (ok)
                ans = max(ans, j - i + 1);
        }
    }

    cout << ans << endl;
}