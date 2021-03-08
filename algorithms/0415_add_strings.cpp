class Solution
{
public:
    string addStrings(string num1, string num2)
    {

        int carry = 0;
        auto num1_size = num1.size();
        auto num2_size = num2.size();
        auto size = num1_size > num2_size ? num1_size : num2_size;
        string ans = "";

        for (std::string::size_type i = 1; i <= size; ++i)
        {
            int n1 = (num1_size - i >= 0) ? (num1[num1_size - i] - '0') : 0;
            int n2 = (num2_size - i >= 0) ? (num2[num2_size - i] - '0') : 0;
            ans = to_string((n1 + n2 + carry) % 10) + ans;
            carry = (n1 + n2 + carry) / 10;
        };

        if (carry > 0)
        {
            ans = "1" + ans;
        }

        return ans;
    }
};