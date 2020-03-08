func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    
    y := len(obstacleGrid)
    x := len(obstacleGrid[0])
    total := uniquePaths(x, y)
    
    for i := 0; i < y; i++ {
        for j := 0; j < x; j++ {
            if obstacleGrid[i][j] == 1 {                                        
                total -= uniquePaths(j+1, i+1) * uniquePaths(x-j, y-i)
            }
        }
    }    
    
    if total < 0 {
        return 0
    }
    return total
}

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
