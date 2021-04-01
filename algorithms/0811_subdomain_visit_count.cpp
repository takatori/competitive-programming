using namespace std;

class Solution
{
    unordered_map<string, int> map;

public:
    vector<string> subdomainVisits(vector<string> &cpdomains)
    {
        for (int i = 0; i < cpdomains.size(); i++)
        {
            vector<string> cpdomain = split(cpdomains[i], ' ');
            int count = stoi(cpdomain[0]);
            string domain = '.' + cpdomain[1];
            for (int j = 0; j < domain.size(); j++)
            {
                if (domain[j] == '.')
                {
                    string subdomain = domain.substr(j + 1);
                    map[subdomain] += count;
                }
            }
        }

        vector<string> ans;
        for (auto &&e : map)
        {
            ans.push_back(to_string(e.second) + " " + e.first);
        }

        return ans;
    }

    vector<string> split(string &s, char delim)
    {
        vector<string> elems;
        stringstream ss(s);
        string item;
        while (getline(ss, item, delim))
        {
            if (!item.empty())
            {
                elems.push_back(item);
            }
        }
        return elems;
    }
};