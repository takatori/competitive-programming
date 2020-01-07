func generateParenthesis(n int) []string {
    
    if n == 1 {
        return []string{"()"}
    }
    
    tmp := generateParenthesis(n-1)
    set := make(map[string]struct{})
    for _, s := range tmp {
        set["(" + s + ")"] = struct{}{}
        set["()" + s ] = struct{}{}
        set[s + "()"] = struct{}{}
    }
    
    var result []string
    
    for k, _ := range set {
        result = append(result, k)
    }
    
    return result
}
