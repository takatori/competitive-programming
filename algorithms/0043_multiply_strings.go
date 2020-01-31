func multiply(num1 string, num2 string) string {
    
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
    
    result := "0"
    
    for _, p := range pos {
        if !(result == "0" && p == 0) {
            result += strconv.Itoa(p)
        }
    }
    
    return result
}
