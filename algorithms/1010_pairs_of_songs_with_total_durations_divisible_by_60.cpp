class Solution
{
public:
    int numPairsDivisibleBy60(vector<int> &time)
    {
        vector<int> remainders(60);
        int count = 0;
        for (int i = 0; i < time.size(); i++)
        {
            if (time.at(i) % 60 == 0)
            { // check if (a % 60 == 0) && (b %% 60 == 0)
                count += remainders.at(0);
            }
            else
            {
                count += remainders.at(60 - time.at(i) % 60); // check if (a % 60) + (b % 60) == 60
            }
            remainders[time.at(i) % 60]++;
        }
        return count;
    }
};