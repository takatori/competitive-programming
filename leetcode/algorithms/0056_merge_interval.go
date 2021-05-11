func merge(intervals [][]int) [][]int {
    
    if len(intervals) < 2 {
        return intervals
    }
    
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    result := [][]int{}
    
    for _, interval := range intervals {
        lastIndex := len(result)-1
        if len(result) == 0 || result[lastIndex][1] < interval[0] { 
            result = append(result, interval)
        } else {
            if result[lastIndex][1] < interval[1] {
                result[lastIndex][1] = interval[1]
            }
        }
    }
    return result
}
