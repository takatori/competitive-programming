func canJump(nums []int) bool {
    
    if len(nums) <= 1 {
        return true
    }
    
    memo := make([]bool, len(nums))
    l := len(nums)
    
    for i := l-1; i >= 0; i-- {
        if nums[i] == 0 {
            memo[i] = false
        } else if nums[i] >= l-1-i {
            memo[i] = true
        } else {
            memo[i] = test(memo[i+1:i+nums[i]+1])
        }
    }
    
    return memo[0]
}

func test(memo []bool) bool {
    
    for _, v := range memo {
        if v {
            return true
        }
    }
    
    return false
}
