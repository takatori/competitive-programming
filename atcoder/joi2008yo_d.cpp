#include <bits/stdc++.h>
using namespace std;

int main()
{
    int M, N;

    cin >> M;
    vector<int> A_x(M);
    vector<int> A_y(M);
    for (int i = 0; i < M; i++)
    {
        cin >> A_x[i] >> A_y[i];
    }

    cin >> N;
    vector<int> B_x(N);
    vector<int> B_y(N);
    unordered_set<string> B_s;
    for (int i = 0; i < N; i++)
    {
        cin >> B_x[i] >> B_y[i];
        B_s.insert(to_string(B_x[i]) + "_" + to_string(B_y[i]));
    }

    int x = A_x[0];
    int y = A_y[0];

    for (int i = 0; i < N; i++)
    {
        int xd = B_x[i] - x;
        int yd = B_y[i] - y;

        bool ok = true;
        for (int j = 1; j < M; j++)
        {
            string s = to_string(A_x[j] + xd) + "_" + to_string(A_y[j] + yd);
            if (B_s.find(s) == B_s.end())
            {
                ok = false;
                break;
            }
        }

        if (ok)
        {
            cout << xd << " " << yd << endl;
            return 0;
        }
    }
}