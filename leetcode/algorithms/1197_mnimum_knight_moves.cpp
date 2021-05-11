class Solution
{
public:
    int minKnightMoves(int x, int y)
    {
        queue<pair<int, int> > q;
        int ans = 0;
        q.push(make_pair(0, 0));
        while (!q.empty())
        {
            auto p = q.front();
            q.pop();
            if (p.first == x && p.second == y)
                break;
            q.push(make_pair(p.first + 2, p.second + 1));
            q.push(make_pair(p.first + 1, p.second + 2));
            q.push(make_pair(p.first + 2, p.second - 1));
            q.push(make_pair(p.first + 1, p.second - 2));
            q.push(make_pair(p.first - 2, p.second + 1));
            q.push(make_pair(p.first - 1, p.second + 2));
            q.push(make_pair(p.first - 2, p.second - 1));
            q.push(make_pair(p.first - 1, p.second - 2));
            ans++;
        }
        return ans;
    }
};