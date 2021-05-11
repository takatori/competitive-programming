var ans []string

func backtrace(n int, s string, left, right int) {
    if len(s) == 2 * n {
        ans = append(ans, s)
        return 
    }
    if left < n {
        backtrace(n, s+"(", left+1, right)
    }
    if right < left {
        backtrace(n, s+")", left, right+1)
    }
}


func generateParenthesis(n int) []string {
    ans = []string{}
    backtrace(n, "", 0, 0)
    return ans
}
