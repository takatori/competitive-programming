func rotate(matrix [][]int)  {
    
    l := len(matrix)-1
    
    for i := 0; i < l-1; i++ {
        for j := i; j < l; j++ {
         swap(matrix, i, j, j, l)
         swap(matrix, i, j, l, l-j)
         swap(matrix, i, j, l-j, i)
        }
    }
    
}


func swap(matrix [][]int, i0,j0,i1,j1 int) {
    tmp := matrix[i1][j1]
    matrix[i1][j1] = matrix[i0][j0]
    matrix[i0][j0] = tmp
}

