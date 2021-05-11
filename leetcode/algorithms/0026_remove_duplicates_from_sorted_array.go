func removeDuplicates(nums []int) int {
    
    if len(nums) == 0 {
        return 0
    }
    
    var i,j int
    max := nums[len(nums)-1]
    
    for i < len(nums) - 1 {
        if nums[i] == max {
            break
        }        
        for j < len(nums) && nums[j] <= nums[i] {
            j++
        }
        i++
        tmp := nums[j]
        nums[j] = nums[i]
        nums[i] = tmp
    }
    nums = nums[:i+1]
    return len(nums)
}
