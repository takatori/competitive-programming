class TweetCounts
{
private:
    map<int, unordered_map<string, int> > db;

public:
    TweetCounts()
    {
    }

    void recordTweet(string tweetName, int time)
    {
        if (db.find(time) != db.end() && db[time].find(tweetName) != db[time].end())
        {
            db[time][tweetName]++;
        }
        else
        {
            db[time] = {{tweetName, 1}};
        }
    }

    vector<int> getTweetCountsPerFrequency(string freq, string tweetName, int startTime, int endTime)
    {
        int f = 60;
        if (freq == "hour")
            f = 3600;
        else if (freq == "day")
            f = 86400;

        int len = (endTime - startTime) / f;
        vector<int> ans(len);

        for (const auto &[time, value] : db)
        {
            if (value.find(tweetName) != value.end())
            {
                int i = (time - startTime) / f;
                if (i >= 0 && i < len)
                    ans[i]++;
            }
        }
        return ans;
    }
};
