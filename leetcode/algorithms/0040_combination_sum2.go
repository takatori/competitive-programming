var result [][]int

func combinationSum2(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    result = make([][]int, 0)
    backtrace([]int{}, candidates, target)
    return result
}

func backtrace(cur []int, candidates []int, target int) {
    if target == 0 {
        result = append(result, cur)
    } else {
        for i := 0; i < len(candidates); i++ {
            v := candidates[i]
            if v <= target {
                tmp := make([]int, 0, len(cur))
                tmp = append(tmp, cur...)
                tmp = append(tmp, v)
                backtrace(tmp, candidates[i+1:], target-v)
                for i < len(candidates)-1 && v == candidates[i+1] {
                    i++
                }
            } else {
                break
            }
        }
    }
}
