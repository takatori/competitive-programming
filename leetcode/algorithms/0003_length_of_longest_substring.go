func lengthOfLongestSubstring(s string) int {
    
    n := len(s)
    m := make(map[byte]int)
    ans := 0
    i := 0 
    j := 0
       
    for ; j < n; j++ {
        if jj, ok := m[s[j]]; ok {
            if jj > i {
                i = jj
            }
        }
        if ans < j - i + 1 {
            ans = j - i + 1
        }
        m[s[j]] = j + 1
    }
    
    return ans
}
    
