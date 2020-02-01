func isMatch(s string, p string) bool {
    
    var i, j int
    // p_len = 10 s_len =17
    // p = ab*cd?i*de, s = abefcdgiescdfimde
    for i < len(p) && j < len(s) { // i = 9, j = 12
        if p[i] == '*' { // p[9] = e
            for i < len(p) -1 && p[i+1] == '*' { // p[8] = d
                i++
            }
            for j < len(s) && (i == len(p)-1 || s[j] != p[i+1]) { // s[8] = e, p[8] = d
                j++  // 8 -> 11 
            }
        } else if p[i] == '?' {
            j++ // j = 7
        } else {
            if p[i] != s[j] { // p[8] = d, s[11] = d
                return false
            } else {
                j++ // j = 12
            }
        }
        i++  // i = 9
    }
    
    return i == len(p) && j == len(s)
}
