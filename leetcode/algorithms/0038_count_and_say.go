func countAndSay(n int) string {
    
    if n == 1 {
        return "1"
    } else {
        
        prev := countAndSay(n-1)
        var i,j int
        var result string
        
        for i < len(prev) {
            if prev[i] == prev[j] {
                i++
            } else {
                result += fmt.Sprintf("%d%s", i-j, string(prev[j]))
                j = i
            }
        }
        
        result += fmt.Sprintf("%d%s", i-j, string(prev[j]))
        
        return result
    }
}
