class Solution
{
public:
    int compress(vector<char> &chars)
    {
        int ans = 0;
        int i = 0;

        while (i < chars.size())
        {
            char head = chars[i];
            int j = i;
            while (j < chars.size() && head == chars[j])
                j++;
            // 圧縮されない
            if (j - i == 1)
            {
                chars[ans] = head;
                ans++;
            }
            else
            {
                string l = to_string(j - i);
                chars[ans] = head;
                ans++;
                for (int k = 0; k < l.length(); k++)
                {
                    chars[ans] = l[k];
                    ans++;
                }
            }
            i = j;
        }
        return ans;
    }
};