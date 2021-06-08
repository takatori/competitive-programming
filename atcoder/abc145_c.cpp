#include <bits/stdc++.h>
using namespace std;

int num;
double total;

void r(int current, queue<int> &unused, vector<vector<double> > &d)
{
    if (unused.size() == 0)
    {
        num++;
        return;
    }
    else
    {
        int a = unused.front();
        unused.pop();
        total += d[current][a];
        r(a, unused, d);
        return;
    }
}

int main()
{
    int N;
    cin >> N;
    vector<vector<int> > c(N, vector<int>(2));
    for (int i = 0; i < N; i++)
    {
        cin >> c[i][0] >> c[i][1];
    }

    vector<vector<double> > d(N, vector<double>(N));
    for (int i = 0; i < N; i++)
    {
        for (int j = i; j < N; j++)
        {
            auto dist = sqrt(pow(c[i][0] - c[j][0], 2.0) + pow(c[i][1] - c[j][1], 2.0));
            d[i][j] = d[j][i] = dist;
        }
    }

    for (int i = 0; i < N; i++)
    {
        for (int j = 0; j < N; j++)
        {
            cout << d[i][j] << " ";
        }
        cout << endl;
    }

    for (int i = 0; i < N; i++)
    {
        queue<int> unused;
        for (int j = 0; j < N; j++)
        {
            if (i != j)
                unused.push(i);
        }
        r(i, unused, d);
    }

    cout << total / num << endl;

    return 0;
}