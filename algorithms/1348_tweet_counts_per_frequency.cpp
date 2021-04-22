class TweetCounts
{
private:
    unordered_map<string, multiset<int> > all;

public:
    TweetCounts()
    {
    }

    void recordTweet(string tweetName, int time)
    {
        all[tweetName].insert(time);
    }

    vector<int> getTweetCountsPerFrequency(string freq, string tweetName, int startTime, int endTime)
    {
        int f = 60;
        if (freq == "hour")
            f = 3600;
        else if (freq == "day")
            f = 86400;

        int len = ((endTime - startTime) / f) + 1;
        vector<int> ans(len);
        const auto s = all.find(tweetName);
        if (s != all.end())
        {
            for (auto t = s->second.lower_bound(startTime); t != s->second.end() && *t <= endTime; ++t)
            {
                ans[(*t - startTime) / f]++;
            }
        }
        return ans;
    }
};
