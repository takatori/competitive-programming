func singleNumber(nums []int) int {
  
	set := make(map[int]struct{})
	total := 0
	for _, v := range nums {   
	  set[v] = struct{}{}
	  total += v
	}
	
	sum := 0
	for k, _ := range set {
	  sum += k
	}
	
	return (3 * sum - total) / 2
  }
  