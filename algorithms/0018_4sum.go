func fourSum(nums []int, target int) [][]int {
    sort.Ints(nums)
    var result [][]int
    
    for i := 0; i < len(nums)-3; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        for j := i+1; j < len(nums)-2; j++ {
            if j > i+1 && nums[j] == nums[j-1] {
                continue
            }
            t, left, right := target - (nums[i] + nums[j]), j+1, len(nums) - 1            
            for left < right {
                sum := nums[left] + nums[right]
                if t == sum {
                    result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
                    left++
                    right--
                    for left < right && nums[left] == nums[left-1] {
                        left++
                    }
                    for left < right && nums[right] == nums[right+1] {
                        right--
                    }
                } else if sum < t {
                    left++
                } else if sum > t {
                    right--
                }
            }            
        }
    }
    
    return result
    
}
