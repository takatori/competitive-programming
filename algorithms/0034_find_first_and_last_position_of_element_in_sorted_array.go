func searchRange(nums []int, target int) []int {
    
    if len(nums) == 0 {
        return []int{-1, -1}
    }
    
    l := 0
    r := len(nums) - 1
    
    for l <= r {
        mid := (l + r) / 2
        if nums[mid] == target {
            start := mid
            end := mid
            for start > 0 && nums[start-1] == target {
                start--
            }
            for end < len(nums)-1 && nums[end+1] == target {
                end++
            }
            return []int{start, end}
        } else if nums[mid] < target {
            l = mid+1
        } else {
            r = mid-1
        }
    }
    
    return []int{-1, -1}
}
