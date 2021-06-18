func combine(n int, k int) [][]int {
 
    result := make([][]int, 0)
    if n < k {
        return result
    }
    
    cur := make([]int, k)
    i := 0    
    for i >= 0 {
        cur[i]++
        if cur[i] > n {
            i--
        } else if i == k-1 {
            dist := make([]int, k)
            copy(dist, cur)
            result = append(result, dist)
        } else {
            i++
            cur[i] = cur[i-1]
        }
    }
    
    return result

}
