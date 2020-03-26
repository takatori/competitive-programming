func subsets(nums []int) [][]int {
    
    if len(nums) == 0 {
        return [][]int{nums}
    }
    
    results := make([][]int, 0)

    prev_results := subsets(nums[1:]) // [3]
    
    for _, subset := range prev_results { // [[], [3]]
        results = append(results, append(subset, nums[0])) // [[2], [2,3]]
    }
    
    return append(prev_results, results...)
}

// []

// [], [3]

// [], [2], [3], [2,3]

// [], [1], [2], [3], [1,2] [1,3], [2,3], [1,2,3]
