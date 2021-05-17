#include <bits/stdc++.h>
using namespace std;

int main()
{
    int N, M;
    cin >> N >> M;

    vector<int> h(M * M, 0);

    for (int i = 0; i < N; i++)
    {
        vector<int> A(M);
        for (int j = 0; j < M; j++)
        {
            cin >> A[j];
        }
        for (int a = 0; a < M - 1; a++)
        {
            for (int b = a + 1; b < M; b++)
            {
                h[a * M + b] += max(A[a], A[b]);
            }
        }
    }
    std::sort(h.begin(), h.end(), std::greater<int>());

    cout << h[0] << endl;
}