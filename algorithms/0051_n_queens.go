var result [][]int
var n int

func solveNQueens(o int) [][]string {
    n = o
    result = [][]int{}
    q,dif,sum := []int{}, []int{}, []int{}
    dfs(q, dif, sum)
    
    ans := make([][]string, 0)
    
    for _, solution := range result {
        tmp := make([]string, o)
        for i, q := range solution {
            tmp[i] = strings.Repeat(".", q) + "Q" + strings.Repeat(".", n-q-1)
        }
        ans = append(ans, tmp)
    }
    
    return ans
}


func dfs(queens, xy_dif, xy_sum []int) {
    p := len(queens)
    if p == n {
        result = append(result, queens)
        return 
    }
    for q := 0; q < n; q++ {
        if !in(queens, q) && !in(xy_dif, p-q) && !in(xy_sum, p+q) {
            dfs(append(queens, q), append(xy_dif, p-q), append(xy_sum, p+q))
        }
    }
}

func in(array []int, target int) bool {
    for _, a := range array {
        if a == target {
            return true
        }
    }
    return false
}
