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
        this->store[key].push_back(make_pair(timestamp, value));
    }

    string get(string key, int timestamp)
    {
        auto it = upper_bound(begin(this->store[key]), end(this->store[key]), pair<int, string>(timestamp, ""), [](
            const pair<int, string>& a, const pair<int, string>& b) { return a.first < b.first; });
        return it == this->store[key].begin() ? "" : prev(it)->second;
    }

};