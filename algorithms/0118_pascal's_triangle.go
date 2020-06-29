func generate(numRows int) [][]int {
  
	if numRows == 0 {
	  return [][]int{}
	}
	
	if numRows == 1 {
	  return [][]int{{1}}
	}
	
	result := [][]int{{1}, {1, 1}}
  
	for i := 2; i < numRows; i++ {
	  tmp := []int{1}
	  for j := 0; j < len(result[i-1])-1; j++ {
		tmp = append(tmp, result[i-1][j]+result[i-1][j+1])
	  }
	  tmp = append(tmp, 1)
	  result = append(result, tmp)
	}
	
	return result
	
}