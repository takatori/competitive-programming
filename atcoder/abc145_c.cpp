#include <bits/stdc++.h>
using namespace std;

int N, X[8], Y[8];

int main()
{
    cin >> N;
    for (int i = 0; i < N; i++)
        cin >> X[i] >> Y[i];

    vector<int> ord;
    for (int i = 0; i < N; i++)
        ord.push_back(i);

    long double sm = 0;
    do
    {
        for (int i = 0; i < N - 1; i++)
        {
            int a = ord[i];
            int b = ord[i + 1];
            long double dx = X[a] - X[b];
            long double dy = Y[a] - Y[b];
            sm += sqrt(dx * dx + dy * dy);
        }
    } while (next_permutation(ord.begin(), ord.end()));

    for (int i = 0; i < N; i++)
        sm /= (i + 1);
    printf("%.10Lf\n", sm);
    return 0;
}