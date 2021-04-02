using namespace std;

class Solution
{
public:
    int jobScheduling(vector<int> &startTime, vector<int> &endTime, vector<int> &profit)
    {
        vector<pair<int, int> > m;
        int max_profit = profit[0];
        m.push_back(make_pair(profit[0], endTime[0]));
        for (int i = 1; i < profit.size(); i++)
        {
            vector<pair<int, int> > tmp;
            if (max_profit < profit[i])
            {
                m.push_back(make_pair(profit[i], endTime[i]));
            }
            for (int j = 0; j < m.size(); j++)
            {
                if (m[j].second <= startTime[i])
                {
                    tmp.push_back(make_pair(m[j].first + profit[i], endTime[i]));
                    max_profit = max(max_profit, m[j].first + profit[i]);
                }
            }
            m = tmp;
        }
        return max_profit;
    }
};