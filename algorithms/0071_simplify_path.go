

func simplifyPath(path string) string {
 
    result := []rune{}
    var prev rune
    for i, c := range path {
        if c == '.' && prev == '.' {
            result = result[:len(result)-2]         
            continue
        } 
        if c != '.' && c != '/' {
            result = append(result, c)
        }
        if c == '/' && prev != '/' && i != len(path)-1 {
            result = append(result, '/')
        }
        prev = c
    }
    
    return string(result)
}

