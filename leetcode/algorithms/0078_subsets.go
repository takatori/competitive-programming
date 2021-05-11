func subsets(nums []int) [][]int {
  
    if len(nums) == 0 {
      return [][]int{nums}
    }
    
    result := make([][]int, 0)
    tmp := subsets(nums[1:])
    
    for _, v := range tmp {
      t := make([]int, len(v))
      copy(t, v)
      result = append(result, append(t, nums[0]))
    }
    
                    
    return append(result, tmp...)
    
  }