using namespace std;

class MedianFinder
{
private:
    list<int> li;
    list<int>::iterator it = this->li.end();

public:
    /** initialize your data structure here. */
    MedianFinder()
    {
    }

    void addNum(int num)
    {
        const int n = this->li.size();
        this->li.insert(lower_bound(this->li.begin(), this->li.end(), num), num);
        if (!n)
            this->it = this->li.begin();
        else if (num < *this->it)
            this->it = n & 1 ? this->it : prev(this->it);
        else
            this->it = n & 1 ? next(this->it) : this->it;
    }

    double findMedian()
    {
        const int n = this->li.size();
        double a = (double)*this->it;
        double b = (double)*next(this->it, n % 2 - 1);
        return (a + b) * 0.5;
    }
};
