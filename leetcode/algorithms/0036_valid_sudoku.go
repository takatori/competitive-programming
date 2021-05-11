func isValidSudoku(board [][]byte) bool {

   for j := 0; j < 9; j++ {
        if rowCheck(board, j) == false {
            return false
        }
    }
    
    for i := 0; i < 9; i++ {
        if columnCheck(board, i) == false {
            return false
        }
    }
    

    for i := 0; i < 9; i += 3 {
        for j := 0; j < 9; j += 3  {
            if boxCheck(board, i, j) == false {
                return false
            }
        }
    } 
    
    return true
}

func boxCheck(board [][]byte, i,j int) bool {
    set := make(map[byte]struct{}, 9)
        
    for ix := i; ix < i+3; ix++ {
        for jx := j; jx < j+3; jx++ {
            n := board[jx][ix]
            if _, ok := set[n]; ok && n != '.' {
                return false
            } else {
                set[n] = struct{}{}
            }            
        }
    }
    
    return true
}

func rowCheck(board [][]byte, j int) bool {
    set := make(map[byte]struct{}, 9)
    
    for i := 0; i < 9; i++ {
        n := board[j][i]
        if _, ok := set[n]; ok && n != '.' {
            return false
        } else {
            set[n] = struct{}{}
        }
    }
    
    return true
}

func columnCheck(board [][]byte, i int) bool {
    
    set := make(map[byte]struct{}, 9)
    
    for j := 0; j < 9; j++ {
        n := board[j][i]
        if _, ok := set[n]; ok && n != '.' {
            return false
        } else {
            set[n] = struct{}{}
        }
    }
    
    return true
}
