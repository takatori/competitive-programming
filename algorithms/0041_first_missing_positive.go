func firstMissingPositive(nums []int) int {
    
    n := len(nums)
    
    for i := 0; i < n; i++ {
        for nums[i] > 0 && nums[i] <= len(nums) && nums[nums[i]-1] != nums[i] {
            tmp := nums[nums[i] - 1]
            nums[nums[i] - 1] = nums[i]
            nums[i] = tmp 
        }
    }
    
    for i := 0; i < n; i++ {
        if nums[i] != i+1 {
            return i+1
        }
    }
    
	return n+1
}
