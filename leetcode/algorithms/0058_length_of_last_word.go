func lengthOfLastWord(s string) int {
    
    var result int
    s = strings.TrimSpace(s)
    
    for i := len(s)-1; i >= 0; i-- {
        if s[i] == ' ' {
            break
        }
        result++
    }
    
    return result
}
