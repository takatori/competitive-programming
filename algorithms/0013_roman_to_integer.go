func romanToInt(s string) int {
    
    dict := map[byte]int {
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }
    
    var result int
    var i int
    
    for i < len(s) {
        if i < len(s) - 1 {
            if dict[s[i]] < dict[s[i+1]] {
                result += (dict[s[i+1]] - dict[s[i]])
                i += 2
                continue
            } 
        }        
        result += dict[s[i]]
        i++
    }
    
    return result    
}
