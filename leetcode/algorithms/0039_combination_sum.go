var result [][]int

func combinationSum(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    result = make([][]int, 0)
    backtrace([]int{}, candidates, target)
    return result
}

func backtrace(cur []int, candidates []int, target int) {
    if target == 0 {
        result = append(result, cur)
    } else {
        for i, v := range candidates {
            if v <= target {
                tmp := make([]int, 0, len(cur))
                tmp = append(tmp, cur...)
                tmp = append(tmp, v)
                backtrace(tmp, candidates[i:], target-v)
            } else {
                break
            }
        }
    }
}
