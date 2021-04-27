using namespace std;

class Solution
{
    int sum;
    int l;
    mt19937 &mt;
    uniform_int_distribution<int> dist;

public:
    Solution(vector<int> &w)
    {
        for (int i : w)
        {
            sum += i;
        }
        l = w.size();
        mt = mt19937 { std::random_device{}(); };
        dist = std::uniform_int_distribution<int>(1, sum);
    }

    int pickIndex()
    {
        int i = (dist(mt) / l) - 1;
        return a[i];
    }
};
