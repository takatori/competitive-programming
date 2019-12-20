func longestCommonPrefix(strs []string) string {
    
    if len(strs) == 0 {
        return ""
    } else if len(strs) == 1 {
        return strs[0]
    }
    
    result := strs[0]
        
    for _, str := range strs[1:] {
        var i int
        if len(str) < len(result) {
            result = result[:len(str)]
        }
        for i < len(result) && i < len(str) {
            if result[i] == str[i] {
                i++
            } else {
                result = result[:i]
                break
            }
        }
        
    }
    
    return result        
}
