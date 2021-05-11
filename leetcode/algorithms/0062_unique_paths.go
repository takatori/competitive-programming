func uniquePaths(m int, n int) int {
    dp := make([]int, m)
    dp[0] = 1
    for i := 0; i < n; i++ {
        for j := 1; j < len(dp); j++ {
            dp[j] += dp[j-1]
        }
    }
    return dp[len(dp)-1]    
}
