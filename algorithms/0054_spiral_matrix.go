func spiralOrder(matrix [][]int) []int {
    
    if len(matrix) == 0 {
        return []int{}
    }
    
    result := make([]int, 0)
    i_len := len(matrix)
    j_len := len(matrix)
    
    var c int
    
    if i_len > j_len {
        c = j_len-1
    } else {
        c = i_len-1
    }
    
    for n := 0; n < c; n++ {
        result = append(result, oneLoop(matrix, n, n))
    }

}

func oneLoop(matrix [][]int, i,j int) []int {
    
    var result []int
    
    for t_j := j; t_j < j_len-j; t_j++ {
        result = append(result, matrix[i][t_j])
    }
    
    for t_i := i+1; t_i < i_len-i-1; t_i++ {
        result = append(result, matrix[t_i][j_len-i-1])
    }
        
    for 
    
    return result
}
