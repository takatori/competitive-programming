using namespace std;

class Solution {  
public:
    int calculate(string s) {
      int ans = 0;
      stack<int> st;
      stack<char> op;
      int i = 0;
      
      s.erase(remove(s.begin(), s.end(),' '), s.end());
      
      while (i < s.length()) {
        
        // 数字の場合
        if (s[i] >= '0' && s[i] <= '9') {
          if (st.empty() || st.size() == op.size()) {
            st.push(s[i] - '0');
          } else {
            int t = st.top();
            st.pop();
            st.push(10 * t + (s[i] - '0'));
          }
        }
        
        // 演算子の場合
        else if (s[i] == '-' || s[i] == '+') {
          if (s[i+1] == '(') {
            op.push(s[i]);
          } else {
            int a = st.top();
            st.pop();
            if (s[i] == '+') {
              st.push(a+(s[i+1] - '0'));
            } else {
              st.push(a-(s[i+1] - '0'));
            }
          }
          i++;
        }
        
        // 括弧が閉じられた場合
        else if (s[i] == ')' && op.size() > 0) {
          int b = st.top();
          st.pop();
          int a = st.top();
          st.pop();
          char o = op.top();
          op.pop();
          if (o == '+') {
            st.push(a+b);
          } else {
            st.push(a-b);
          }
        }
        
        i++;
      }
      
      return st.top();
    }
};