class TimeMap
{
private:
    unordered_map<string, vector<pair<int, string>>> store;

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
            this->store[key] = vector<pair<int, string>>{newItem};
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
        auto items = this->store[key];
        int idx = upper_bound(items, timestamp);
        if (idx == -1)
        {
            return "";
        }
        return items.at(idx - 1).second;
    }

    int upper_bound(vector<pair<int, string>> arr, int timestamp)
    {
        int mid;
        int left = 0;
        int right = arr.size();

        while (left < right)
        {
            mid = (right - left) / 2;
            if (timestamp >= arr.at(mid).first)
            {
                left = mid + 1;
            }
            else
            {
                right = mid;
            }
        }

        if (right == 0)
        {
            return -1;
        }

        return left;
    }
};