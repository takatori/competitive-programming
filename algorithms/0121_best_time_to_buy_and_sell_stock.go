
func maxProfit(prices []int) int {
    
	if len(prices) == 0 {
	  return 0
	}
	
	var max int
  
	for i := 0; i < len(prices)-1; i++ {
	  for j := i+1; j < len(prices); j++ {
		diff := prices[j] - prices[i]
		if diff > max {
		  max = diff
		}
	  }
	}
	
	return max
	  
  }
  
  
  