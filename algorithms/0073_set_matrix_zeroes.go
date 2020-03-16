func setZeroes(matrix [][]int)  {
    
    i_s := make(map[int]struct{}, 0)
    j_s := make(map[int]struct{}, 0)
    
    i_l := len(matrix)
    j_l := len(matrix[0])
    
    for i := 0; i < i_l; i++ {
        for j := 0; j < j_l; j++ {
            if matrix[i][j] == 0 {
                i_s[i] = struct{}{}
                j_s[j] = struct{}{}
            }
        }
    }
    
    for i, _ := range i_s {
        for j := 0; j < j_l; j++ {
            matrix[i][j] = 0
        }
    }
    
    for j, _ := range j_s {
        for i := 0; i < i_l; i++ {
            matrix[i][j] = 0
        }
    }
}
