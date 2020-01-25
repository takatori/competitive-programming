var choices = [...]byte{'1','2','3','4','5','6','7','8','9'}

func solveSudoku(board [][]byte)  {
    choiceMaker(board, 0, 0)
}

func choiceMaker(board [][]byte, row, column int) bool {
    if row == 9 {
        return true
    }
    // next spot
    nextRow, nextColumn := row, column + 1
    if nextColumn == 9 {
        nextRow, nextColumn = row+1, 0
    }
    // fixed
    if board[row][column] != '.' {
        return choiceMaker(board, nextRow, nextColumn)
    }
    // try all possibilities!
    for _, option := range choices {
        if validChoice(board, option, row, column) {
            // place choice
            board[row][column] = option
            if !choiceMaker(board, nextRow, nextColumn) {
                // remove
                board[row][column] = '.'
            } else {
                return true
            }
        }
    }
    return false
}

func validChoice(board [][]byte, choice byte, row, column int) bool {
    // row & col
    for i := 0; i < 9; i++ {
        if board[row][i] == choice || board[i][column] == choice {
            return false
        }
    }
    // square
    rowLock, colLock := 0, 0
    if 2 < column && column < 6 { // mid
        colLock = 3
    } else if column > 5 { // right
        colLock = 6
    }
    if 2 < row && row < 6 { // mid
        rowLock = 3
    } else if row > 5 { // bot
        rowLock = 6
    }
    for i := 0; i < 3; i++ {
        if (board[rowLock+i][colLock] == choice || 
            board[rowLock+i][colLock+1] == choice || 
            board[rowLock+i][colLock+2] == choice) {
            return false
        }
    }
    return true
}
