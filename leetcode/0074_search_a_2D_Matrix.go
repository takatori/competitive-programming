func searchMatrix(matrix [][]int, target int) bool {
 
    if len(matrix) == 0 {
        return false
    }
    
    m := len(matrix)
    n := len(matrix[0])
    
    row := -1
    
    for i := 0; i < m-1; i++ {
        currentRow := matrix[i][0]
        peek := matrix[i+1][0]
        if target >= currentRow && target < peek {
            row = i
        }
    }
    
    if row == -1 {
        row = m-1
    }
    
    for j := 0; j < n; j++ {
        if matrix[row][j] == target {
            return true
        }
    }
    
    return false
    
}
