using namespace std;

class Solution
{
public:
    vector<vector<int> > kClosest(vector<vector<int> > &points, int k)
    {
        auto N = points.size();
        vector<pair<long, int> > dists(N);

        for (int i = 0; i < N; i++)
        {
            auto point = points.at(i);
            long dist = point[0] * point[0] + point[1] * point[1];
            dists.at(i) = make_pair(dist, i);
        }

        std::sort(dists.begin(), dists.end(), [](pair<long, int> a, pair<long, int> b) {
            return a.first < b.first;
        });

        vector<vector<int> > result(k);

        for (int i = 0; i < k; i++)
        {
            result.at(i) = points.at(dists.at(i).second);
        }

        return result;
    }
};