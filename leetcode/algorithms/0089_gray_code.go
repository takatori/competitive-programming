func grayCode(n int) []int {
	if n < 0 {
	  return nil
	}
	
	result := []int{0}
	if n == 0 {
	  return result
	}
  
	for i := 0; i < n; i++ {
	  for j := len(result) - 1; j >= 0; j-- {
		N := 1 << i
		result = append(result, (N | result[j]))
	  }
	}
	
	return result    
  }