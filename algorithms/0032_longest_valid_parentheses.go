var memo map[int]int

func longestValidParentheses(s string) int {
    
    i := len(s) - 2 
    memo = make(map[int]int, len(s))
    var result int
    
    for i >= 0 {
        if s[i] == ')' {
        } else {
            length := calc(s, i) 
            memo[i] = length
            if length > result {
                result = length
            }
        }
        i--
    }
    return result
}

func calc(s string, i int) int {
    
    if s[i+1] == ')' {
        if l, ok := memo[i+2]; ok {
            return l + 2
        } else {
            return 2
        }
    }
    
    if s[i+1] == '(' {
        if l, ok := memo[i+1]; ok {
            if i+l+1 < len(s) && s[i+l+1] == ')' {
                return l + 2 + memo[i+l+2]
            } else {
                return 0
            }
        } else {
            return 0
        }
    }    
    
    return 0
}
