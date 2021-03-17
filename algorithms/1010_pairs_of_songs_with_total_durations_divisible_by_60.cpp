#include<bits/stdc++.h>

using namespace std;

class Solution
{
public:
    int numPairsDivisibleBy60(vector<int> &time)
    {
        int N = time.size();
        vector<vector<int>> memo(N);
        for (int i = 0; i < N; i++)
        {
            memo.at(i) = vector<int>(N);
        }
        int ans = 0;
        for (int i = 0; i < time.size() - 1; i++)
        {
            for (int j = i + 1; j < time.size(); j++)
            {
                if ((time.at(i) + time.at(j)) % 60 == 0)
                    ans++;
            }
        }
        return ans;
    }
};
