class TimeMap
{
private:
    unordered_map<string, vector<pair<int, string> > > store;

public:
    /** Initialize your data structure here. */
    TimeMap()
    {
    }

    void set(string key, string value, int timestamp)
    {
        auto newItem = make_pair(timestamp, value);
        if (this->store.count(key) == 0)
        {
            this->store[key] = vector<pair<int, string> >{newItem};
        }
        else
        {
            this->store[key].push_back(newItem);
        }
    }

    string get(string key, int timestamp)
    {
        if (this->store.count(key) == 0)
        {
            return "";
        }
        string ans = "";
        auto items = this->store[key];
        for (int i = 0; i < items.size(); i++)
        {
            if (items.at(i).first <= timestamp)
            {
                ans = items.at(i).second;
            }
        }
        return ans;
    }
};