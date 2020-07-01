func getRow(rowIndex int) []int {
    
	if rowIndex == 0 {
	  return []int{1}
	}
	
  result := []int{1,1}
  
	for i := 1; i < rowIndex; i++ {
	  tmp := []int{1}
	  for j := 0; j < len(result)-1; j++ {
  		tmp = append(tmp, result[j]+result[j+1])
	  }
	  tmp = append(tmp, 1)
	  result = tmp
	}
	
	return result
}