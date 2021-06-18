func singleNumber(nums []int) int {
    
	if len(nums) == 1 {
	  return nums[0]
	}
	
	sort.Ints(nums)
	
	for i := 1; i < len(nums)-1; i++ {
	  if nums[i] != nums[i-1] && nums[i] != nums[i+1] {
		return nums[i]
	  }
	}
	
	if nums[0] != nums[1] {
	  return nums[0]   
	} else {
	  return nums[len(nums)-1]
	}
}