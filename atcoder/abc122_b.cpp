#include <bits/stdc++.h>
using namespace std;

int main()
{
    string S;
    cin >> S;
    int ans = 0;
    int i = 0;
    int j = 0;

    while (i < S.length())
    {
        int tmp = 0;
        j = i;
        while (j < S.length())
        {
            if (S[j] == 'A' || S[j] == 'C' || S[j] == 'G' || S[j] == 'T')
            {
                tmp++;
                j++;
            }
            else
            {
                break;
            }
        }
        ans = max(ans, tmp);
        i++;
    }

    cout << ans << endl;
}