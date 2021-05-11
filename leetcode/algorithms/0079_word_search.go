func exist(board [][]byte, word string) bool {
  
	if len(board) == 0 {
	  return false
	}
	
	for i := 0; i < len(board); i++ {
	  for j := 0; j < len(board[0]); j++ {
		if board[i][j] == word[0] {
		  if check(board, word, j, i) {
			return true
		  }
		}
	  }
	}
	return false
  }
  
  
  func check(board [][]byte, word string, x, y int) bool {
	
	if len(word) == 0 {
	  return true
	}
	
	if x < 0 || x >= len(board[0]) || y < 0 || y >= len(board) {
	  return false
	}
	
	var tmp byte
	
	if board[y][x] == word[0] {
	  tmp = board[y][x]
	  board[y][x] = '_'
	  if check(board, word[1:], x, y-1) {
		return true
	  }
	  if check(board, word[1:], x-1, y) {
		return true
	  }
	  if check(board, word[1:], x, y+1) {
		return true
	  }
	  if check(board, word[1:], x+1, y) {
		return true
	  }
	  
	  board[y][x] = tmp
	  return false
	}
	
	return false
  }