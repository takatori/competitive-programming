class Solution
{
public:
    vector<int> topKFrequent(vector<int> &nums, int k)
    {
        sort(nums.begin(), nums.end());

        auto c = [](pair<int, int> l, pair<int, int> r) {
            return l.second < r.second;
        };
        priority_queue<pair<int, int>, vector<pair<int, int> >, decltype(c)> q(c);

        int i = 0;
        while (i < nums.size())
        {
            int cnt = 1;
            int j = i + 1;
            while (j < nums.size() && nums[i] == nums[j])
            {
                cnt++;
                j++;
            }
            q.push(make_pair(nums[i], cnt));
            i = j;
        }

        vector<int> ans;
        for (int i = 0; i < k; i++)
        {
            auto a = q.top();
            q.pop();
            ans.push_back(a.first);
        }
        return ans;
    }
};