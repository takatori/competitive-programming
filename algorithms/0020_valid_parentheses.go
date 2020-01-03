func isValid(s string) bool {
    
    dict := map[rune]rune {
        ')': '(',
        ']': '[',
        '}': '{',
    }
    
    stack := make([]rune, 0)
    
    for _, c := range s {
        
        if len(stack) == 0 {
            stack = append(stack, c)
            continue
        } 

        if top, ok := dict[c]; ok && top == stack[len(stack)-1] {
            stack = stack[:len(stack)-1]
        } else {                        
            stack = append(stack, c)            
        }

    }
    
    return len(stack) == 0

}
