func getPermutation(n int, k int) string {
    
    s := make(map[int]struct{}, n)
    for i := 1; i <= n; i++ {
        s[i] = struct{}{}
    }
    
    result := make([]int, n)
    
    for i := n; i >= 0; i-- {
        result = append(result, f(s, k))
    }
    
    return strings.Join(result, "")
}

func f(s map[int]struct{}, k) int {
    
}

func fuctorial(n int) int {
    if n == 1 {
        return n
    }
    return n * fuctorial(n-1)
}
