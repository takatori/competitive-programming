func longestPalindrome(s string) string {
    
    if s == "" {
        return ""
    }

    length := len(s)
    memo := make([][]bool, length)
    var mi, mj, max int
    
    for i := length - 1; i >= 0; i-- {
        memo[i] = make([]bool, length)
        for j := length - 1;  j >= i; j-- {
            if (i == j) || (j - i == 1 && s[i] == s[j]) {
                memo[i][j] = true
            } else {
                memo[i][j] = s[i] == s[j] && memo[i+1][j-1]
            }
            
            if memo[i][j] && max < j - i {
                max = j - i
                mi = i
                mj = j
            }
        }
    }
    
    return s[mi:mj+1]
}
