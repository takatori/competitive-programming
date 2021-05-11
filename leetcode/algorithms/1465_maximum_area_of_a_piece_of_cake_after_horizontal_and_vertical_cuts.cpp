class Solution
{
public:
    int maxArea(int h, int w, vector<int> &horizontalCuts, vector<int> &verticalCuts)
    {
        horizontalCuts.push_back(0);
        horizontalCuts.push_back(h);
        verticalCuts.push_back(0);
        verticalCuts.push_back(w);
        sort(horizontalCuts.begin(), horizontalCuts.end());
        sort(verticalCuts.begin(), verticalCuts.end());

        long long maxH = 0;
        long long maxW = 0;
        long long m = 1000000000 + 7;

        for (int i = 0; i < horizontalCuts.size() - 1; i++)
        {
            int tmpH = horizontalCuts[i + 1] - horizontalCuts[i];
            if (tmpH > maxH)
                maxH = tmpH;
        }

        for (int j = 0; j < verticalCuts.size() - 1; j++)
        {
            int tmpW = verticalCuts[j + 1] - verticalCuts[j];
            if (tmpW > maxW)
                maxW = tmpW;
        }

        return int((maxH * maxW) % m);
    }
};