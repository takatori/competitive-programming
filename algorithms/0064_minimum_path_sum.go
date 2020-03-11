func minPathSum(grid [][]int) int {
    
    m := len(grid)
    n := len(grid[0])
    
    for i := 1; i < m; i++ {
        grid[i][0] = grid[i][0] + grid[i-1][0]
    }
    
    for j := 1; j < n; j++ {
        grid[0][j] = grid[0][j] + grid[0][j-1]
    }
    
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            var min int
            top := grid[i-1][j]
            left := grid[i][j-1]
            if top < left {
                min = top
            } else {
                min = left
            }
            grid[i][j] = grid[i][j] + min
        }
    }
    
    return grid[m-1][n-1]
}
