class Solution
{
public:
    int subarraySum(vector<int> &nums, int k)
    {
        int i = 0;
        int ans = 0;
        for (int i = 0; i < nums.size(); i++)
        {
            int current = 0;
            for (int j = i; j < nums.size(); j++)
            {
                current += nums.at(j);
                if (current == k)
                {
                    ans++;
                }
            }
        }
        return ans;
    }
};