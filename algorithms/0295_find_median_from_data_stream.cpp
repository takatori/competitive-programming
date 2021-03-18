#include <list>

using namespace std;

class MedianFinder
{
private:
    list<int> li;
    list<int>::iterator it = li.begin();

public:
    /** initialize your data structure here. */
    MedianFinder()
    {
    }

    void addNum(int num)
    {
        this->li.insert(lower_bound(this->li.begin(), this->li.end(), num), num);
        if (*this->it < num && next(this->it) != this->li.end())
        {
            this->it = next(this->it);
        }
        else if (*this->it >= num && prev(this->it) != this->li.begin())
        {
            this->it = prev(this->it);
        }
    }

    double findMedian()
    {
        if (this->li.size() % 2 == 0)
        {
            return double(*this->it + *next(this->it)) / 2;
        }
        else
        {
            return double(*this->it);
        }
    }
};
