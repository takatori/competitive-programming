func insert(intervals [][]int, newInterval []int) [][]int {
    
    if len(intervals) == 0 {
        return [][]int{newInterval}
    }
    
    if len(newInterval) == 0 {
        return intervals
    }
    
    result := [][]int{}
    i := 0
    
    for ; i < len(intervals); i++ {
        l := len(result)-1
        interval := intervals[i]
        if interval[0] < newInterval[0] {
            result = append(result, interval)
        } else {
            if result[l][1] < newInterval[0] {
                result = append(result, newInterval)
            } else if result[l][1] < newInterval[1] {
                result[l][1] = newInterval[1]
            }
            break
        }
    }
    
    for j := i; j < len(intervals); j++ {
        l := len(result)-1
        interval := intervals[j]
        if len(result) == 0 || result[l][1] < interval[0] { 
            result = append(result, interval)
        } else {
            if result[l][1] < interval[1] {
                result[l][1] = interval[1]
            }
        }
    }
    return result    
}
