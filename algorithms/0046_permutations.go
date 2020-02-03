var memo map[string][][]int

func permute(nums []int) [][]int {
    
    sort.Ints(nums) // [1,2,3]
    key := fmt.Sprint(nums) // 1,2,3
    if memo == nil {
        memo = make(map[string][][]int)
    }
    if len(nums) <= 1 {
        memo[key] = [][]int{nums}
        return [][]int{nums}
    }   
    if n, ok := memo[key]; ok {
        return n
    }
    
    result := make([][]int, 0)
    
    for i := 0; i < len(nums); i++ { // i = 1
        target := append(nums[:i], nums[i+1:]...) // [1], [3] => [1,3]
        for _, prev := range permute(target) { // [[1,3],[3,1]]
            result = append(result, append([]int{nums[i]}, prev...)) // [1,2,3], [1,3,2],[2,1,3],[2,3,1],
        }
    }
    
    memo[key] = result
    
    return memo[key]
}
