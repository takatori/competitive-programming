func threeSumClosest(nums []int, target int) int {
    
    var result int
    mindiff := math.MaxInt32
    sort.Ints(nums)
    
    for i := 0; i < len(nums) - 2; i++ {
        left, right := i+1, len(nums) - 1
        for left < right {
            sum := nums[i] + nums[left] + nums[right]
            diff := abs(target - sum)
            if diff < mindiff {
                mindiff = diff
                result = sum               
            }
            if target - sum >= 0 {
                left++
            } else if target - sum < 0 {
                right--
            }
        }
    }
    return result
}

func abs(x int) int {
    if x < 0 {
        return -x
    } else {
        return x
    }
}
