using namespace std;

class Solution
{
public:
    string minRemoveToMakeValid(string s)
    {
        stack<int> st;
        for (int i = 0; i < s.size(); i++)
        {
            auto c = s.at(i);
            if (st.empty() && c == ')')
                s[i] = '-';
            else if (c == '(')
                st.push(i);
            else if (c == ')')
                st.pop();
        }

        while (!st.empty())
        {
            s[st.top()] = '-';
            st.pop();
        }

        string result;
        for (int i = 0; i < s.size(); i++)
        {
            if (s[i] != '-')
                result += s[i];
        }

        return result;
    }
};