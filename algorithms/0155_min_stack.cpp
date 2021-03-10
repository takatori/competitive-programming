class MinStack
{
public:
    /** initialize your data structure here. */
    MinStack()
    {
        this->minValueStack.push_back(2147483647);
    }

    void push(int x)
    {
        this->stack.push_back(x);
        int min = this->minValueStack.back();
        if (x < min)
        {
            this->minValueStack.push_back(x);
        }
        else
        {
            this->minValueStack.push_back(min);
        }
    }

    void pop()
    {
        this->stack.pop_back();
        this->minValueStack.pop_back();
    }

    int top()
    {
        return this->stack.back();
    }

    int getMin()
    {
        return this->minValueStack.back();
    }

private:
    std::vector<int> stack;
    std::vector<int> minValueStack;
};

/**
 * Your MinStack object will be instantiated and called as such:
 * MinStack* obj = new MinStack();
 * obj->push(x);
 * obj->pop();
 * int param_3 = obj->top();
 * int param_4 = obj->getMin();
 */