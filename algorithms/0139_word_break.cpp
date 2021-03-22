using namespace std;

class Solution
{
public:
    bool wordBreak(string s, vector<string> &wordDict)
    {
        if (s.size() == 0)
            return true;
        for (string word : wordDict)
        {
            if (s.size() >= word.size() && equal(begin(word), end(word), begin(s)))
                if (wordBreak(s.substr(word.size()), wordDict))
                    return true;
        }
        return false;
    }
};