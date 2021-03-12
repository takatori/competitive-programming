using namespace std;

class RandomizedSet
{
private:
    unordered_map<int, int> dict;
    vector<int> list;

public:
    /** Initialize your data structure here. */
    RandomizedSet()
    {
    }

    /** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
    bool insert(int val)
    {
        if (this->dict.find(val) != this->dict.end())
        {
            return false;
        }
        this->list.push_back(val);
        this->dict[val] = this->list.size() - 1;
        return true;
    }

    /** Removes a value from the set. Returns true if the set contained the specified element. */
    bool remove(int val)
    {
        if (this->dict.find(val) == this->dict.end())
        {
            return false;
        }
        int idx = this->dict[val];
        this->list.at(idx) = this->list.at(this->list.size() - 1);
        this->list.pop_back();
        this->dict.erase(val);
        return true;
    }

    /** Get a random element from the set. */
    int getRandom()
    {
        mt19937 mt{std::random_device{}()};
        uniform_int_distribution<int> dist(0, this->list.size() - 1);
        return this->list.at(dist(mt));
    }
};
