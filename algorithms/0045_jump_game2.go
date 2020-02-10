func jump(nums []int) int {
    
    memo := make([]int, len(nums))
    last := len(nums)-1
    
    for i := last-1; i >= 0; i-- {
        
        distance := last-i
        jump := nums[i]
        
        if jump >= distance {
            memo[i] = 1
        } else {
            min := math.MaxInt32
            for j := 1; j <= jump; j++ {
                if memo[i+j] < min {
                    min = memo[i+j]
                }
            }
            memo[i] = min+1
        }
    }
    
    return memo[0]
}
