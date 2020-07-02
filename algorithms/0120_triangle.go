func minimumTotal(triangle [][]int) int {
    
	if len(triangle) == 0 {
	  return 0
	}
	
	l := len(triangle)
	
	return total(triangle, 0, 0, l)
  }
  
  func total(triangle [][]int, i, j, l int) int {
	
	t := triangle[i][j]
	
	if l-1 == i {
	  return t
	}
	
	left_total := total(triangle, i+1, j, l)
	right_total := total(triangle, i+1, j+1, l)
	
	if left_total < right_total {
	  return left_total + t
	} else {
	  return right_total + t
	}
  }