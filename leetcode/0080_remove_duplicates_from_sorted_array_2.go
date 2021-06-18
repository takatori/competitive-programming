func removeDuplicates(nums []int) int {
 
	i := 0
	next := 0
	total := 0
   
	for i < len(nums) {
	  c := 0
	  for nums[i] == nums[next] {
		if c > 1 {
		  nums[next] = math.MaxInt32
		} else {
		  total++
		}
		next++
		c++
		if next >= len(nums) {
		  break
		}
	  }
	  i = next
	} 
	
	sort.Ints(nums)
	return total
  }