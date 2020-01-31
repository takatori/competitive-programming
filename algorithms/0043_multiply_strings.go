func multiply(num1 string, num2 string) string {
    
    if num1 == "0" || num2 == "0" {
        return "0"
    }
    
    m := len(num1)
    n := len(num2)
    
    pos := make([]int, m+n)
    
    for i := m-1; i >= 0; i-- {
        for j := n-1; j >= 0; j-- {
            mul := int((num1[i] - '0') * (num2[j] - '0')) 
            p1 := i+j 
            p2 := i+j+1 
            sum := mul + pos[p2] 
            pos[p1] += sum / 10 
            pos[p2] = sum % 10  
        }
    }
    
    result := ""
    
    for _, p := range pos {
        if !(result == "" && p == 0) {
            result += strconv.Itoa(p)
        }
    }
    
    return result
}
