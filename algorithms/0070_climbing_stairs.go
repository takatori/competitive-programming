var memo map[int]int

func climbStairs(n int) int {
    if memo == nil {
        memo = make(map[int]int, 0)
    }
    if v, ok := memo[n]; ok {
        return v
    }
    if n < 3 {
        memo[n] = n
        return n
    }
    s := climbStairs(n-1) + climbStairs(n-2)
    memo[n] = s
    return s
}
