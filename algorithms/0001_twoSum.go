func twoSum(nums []int, target int) []int {    
    
    memo := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        if j, ok := memo[nums[i]]; ok {
            return []int{j, i}
        } else {
            memo[target - nums[i]] = i
        }
    }
    return nil
}
