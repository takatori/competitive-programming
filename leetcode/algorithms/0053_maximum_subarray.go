func maxSubArray(nums []int) int {
    
    max := math.MinInt32
    
    for i := 0; i < len(nums); i++ {
        sum := 0
        for j := i; j < len(nums); j++ {
            sum += nums[j]
            if sum > max {
                max = sum
            }
        }
    }
    
    return max
}
