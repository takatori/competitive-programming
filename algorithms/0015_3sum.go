func threeSum(nums []int) [][]int {
    
    memo := make(map[string][]int, 0)
    
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            for k := j + 1; k < len(nums); k++ {
                if nums[i] + nums[j] + nums[k] == 0 {                                      
                    l := []int{nums[i], nums[j], nums[k]}
                    sort.Ints(l)
                    memo[fmt.Sprint(l)] = l
                }
            }
        }
    }
    
    var result [][]int
    for _, v := range memo {
        result = append(result, v)
    }
    return result
}

