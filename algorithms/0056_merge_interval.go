func merge(intervals [][]int) [][]int {
    
    if len(intervals) < 2 {
        return intervals
    }
    
    result := [][]int{intervals[0]}
    for i := 1; i < len(intervals); i++ {
        lastIdx := len(result)-1
        first, second := mergeTwo(result[lastIdx], intervals[i])
        if second == nil {
            result[lastIdx] = first
        } else {
            result[lastIdx] = first
            result = append(result, second)
        }
    }
    return result
}

func mergeTwo(left []int, right []int) ([]int, []int) {
    
    if left[0] > right[1] {
        return right, left
    }
    
    if left[0] <= right[0] && right[1] <= left[1] {
        return left, nil
    }
    
    if right[0] <= left[0] && left[1] <= right[1] {
        return right, nil
    }
    
    if left[1] >= right[0] {
        l := left[0]
        r := left[1]
        if left[0] > right[0] {
            l = right[0]
        }
        if left[1] < right[1] {
            r = right[1]
        }
        return []int{l,r}, nil
    } 
    
    return left, right
}
